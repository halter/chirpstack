use std::collections::HashMap;
use std::sync::RwLock;
use std::time::Instant;

use anyhow::Result;
use tracing::{error, info};

use crate::monitoring::prometheus;
use diesel::{ConnectionError, ConnectionResult};
use diesel_async::pooled_connection::deadpool::{Object as DeadpoolObject, Pool as DeadpoolPool};
use diesel_async::pooled_connection::{AsyncDieselConnectionManager, ManagerConfig};
use diesel_async::{AsyncConnection, AsyncPgConnection};
use futures::{future::BoxFuture, FutureExt};
use prometheus_client::metrics::histogram::{exponential_buckets, Histogram};
use scoped_futures::ScopedBoxFuture;

use crate::config;

use crate::helpers::tls::get_root_certs;

pub type AsyncPgPool = DeadpoolPool<AsyncPgConnection>;
pub type AsyncPgPoolConnection = DeadpoolObject<AsyncPgConnection>;

lazy_static! {
    // Changed from single pool to HashMap of pools
    static ref ASYNC_PG_POOLS: RwLock<HashMap<String, AsyncPgPool>> = RwLock::new(HashMap::new());
    static ref STORAGE_PG_CONN_GET: Histogram = {
        let histogram = Histogram::new(exponential_buckets(0.001, 2.0, 12));
        prometheus::register(
            "storage_pg_conn_get_duration_seconds",
            "Time between requesting a PostgreSQL connection and the connection-pool returning it",
            histogram.clone(),
        );
        histogram
    };
}

// Modified to accept a pool identifier
pub fn setup(pool_id: &str, conf: &config::Postgresql) -> Result<()> {
    info!("Setting up PostgreSQL connection pool for {}", pool_id);
    let mut config = ManagerConfig::default();
    if pool_id == "remote" {
        config.custom_setup = Box::new(pg_establish_connection);
    } else if pool_id == "local" {
        config.custom_setup = Box::new(pg_establish_connection_local);
    } else {
        return Err(anyhow!("Unknown pool identifier: {}", pool_id));
    }
    info!("Setting up DSN for {}: {}", pool_id, conf.dsn);
    let mgr = AsyncDieselConnectionManager::<AsyncPgConnection>::new_with_config(&conf.dsn, config);
    let pool = DeadpoolPool::builder(mgr)
        .max_size(conf.max_open_connections as usize)
        .build()?;
    set_async_db_pool(pool_id, pool);

    Ok(())
}

// Source remains unchanged
fn pg_establish_connection(config: &str) -> BoxFuture<ConnectionResult<AsyncPgConnection>> {
    let fut = async {
        let conf = config::get();

        let root_certs = get_root_certs(if conf.postgresql.ca_cert.is_empty() {
            None
        } else {
            Some(conf.postgresql.ca_cert.clone())
        })
        .map_err(|e| ConnectionError::BadConnection(e.to_string()))?;
        let rustls_config = rustls::ClientConfig::builder()
            .with_root_certificates(root_certs)
            .with_no_client_auth();
        let tls = tokio_postgres_rustls::MakeRustlsConnect::new(rustls_config);
        let (client, conn) = tokio_postgres::connect(config, tls)
            .await
            .map_err(|e| ConnectionError::BadConnection(e.to_string()))?;
        tokio::spawn(async move {
            if let Err(e) = conn.await {
                error!(error = %e, "PostgreSQL connection error");
            }
        });
        AsyncPgConnection::try_from(client).await
    };
    fut.boxed()
}

fn pg_establish_connection_local(config: &str) -> BoxFuture<ConnectionResult<AsyncPgConnection>> {
    let fut = async {
        let conf = config::get();
        let (client, conn) = tokio_postgres::connect(config, tokio_postgres::NoTls)
            .await
            .map_err(|e| ConnectionError::BadConnection(e.to_string()))?;
        tokio::spawn(async move {
            if let Err(e) = conn.await {
                error!(error = %e, "PostgreSQL connection error");
            }
        });
        AsyncPgConnection::try_from(client).await
    };
    fut.boxed()
}

// Modified to accept pool_id
fn get_async_db_pool(pool_id: &str) -> Result<AsyncPgPool> {
    let pools_r = ASYNC_PG_POOLS.read().unwrap();
    let pool = pools_r
        .get(pool_id)
        .ok_or_else(|| {
            anyhow!(
                "PostgreSQL connection pool '{}' is not initialized",
                pool_id
            )
        })?
        .clone();
    Ok(pool)
}

// Modified to accept pool_id
pub async fn get_async_db_conn() -> Result<AsyncPgPoolConnection> {
    return get_async_db_conn_by_id("remote").await;
}

pub async fn get_async_db_conn_by_id(pool_id: &str) -> Result<AsyncPgPoolConnection> {
    let pool = get_async_db_pool(pool_id)?;

    let start = Instant::now();
    let res = pool.get().await?;

    STORAGE_PG_CONN_GET.observe(start.elapsed().as_secs_f64());

    Ok(res)
}

// Transaction function remains largely unchanged, just works with the connection
pub async fn db_transaction<'a, R, E, F>(
    conn: &mut AsyncPgPoolConnection,
    callback: F,
) -> Result<R, E>
where
    F: for<'r> FnOnce(&'r mut AsyncPgPoolConnection) -> ScopedBoxFuture<'a, 'r, Result<R, E>>
        + Send
        + 'a,
    E: From<diesel::result::Error> + Send + 'a,
    R: Send + 'a,
{
    conn.transaction(callback).await
}

// Modified to accept pool_id
fn set_async_db_pool(pool_id: &str, p: AsyncPgPool) {
    let mut pools_w = ASYNC_PG_POOLS.write().unwrap();
    pools_w.insert(pool_id.to_string(), p);
}

// Optional: Add a function to list available pools
pub fn get_available_pools() -> Vec<String> {
    let pools_r = ASYNC_PG_POOLS.read().unwrap();
    pools_r.keys().cloned().collect()
}
