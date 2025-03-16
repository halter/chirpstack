use std::net::SocketAddr;
use std::sync::Arc;
use std::time::Duration;

use anyhow::Result;
use axum::{
    body::Bytes,
    response::{IntoResponse, Json, Response},
    Router,
};
use chrono::Utc;
use http::StatusCode;
use redis::streams::StreamReadReply;
use rustls::{
    server::{NoClientAuth, WebPkiClientVerifier},
    ServerConfig,
};
use serde::Serialize;
use tokio::sync::oneshot;
use tokio::task;
use tracing::{error, info, span, warn, Instrument, Level};

use crate::backend::joinserver;
use crate::helpers::errors::PrintFullError;
use crate::helpers::tls::{get_root_certs, load_cert, load_key};
use crate::storage::{error::Error as StorageError, get_async_redis_conn, redis_key};
use crate::uplink::error::Error as UplinkError;
use crate::{config, stream};
use backend::{BasePayload, BasePayloadResultProvider, MessageType};
use chirpstack_api::stream as stream_pb;
use lrwn::EUI64;

pub async fn setup() -> Result<()> {
    let conf = config::get();
    if conf.backend_interfaces.bind.is_empty() {
        info!("Backend interfaces API interface is disabled");
        return Ok(());
    }

    let addr: SocketAddr = conf.backend_interfaces.bind.parse()?;
    info!(bind = %conf.backend_interfaces.bind, "Setting up backend interfaces API");

    let app = Router::new().fallback(handle_request);

    if !conf.backend_interfaces.ca_cert.is_empty()
        || !conf.backend_interfaces.tls_cert.is_empty()
        || !conf.backend_interfaces.tls_key.is_empty()
    {
        let mut server_config = ServerConfig::builder()
            .with_client_cert_verifier(if conf.backend_interfaces.ca_cert.is_empty() {
                Arc::new(NoClientAuth)
            } else {
                let root_certs = get_root_certs(Some(conf.backend_interfaces.ca_cert.clone()))?;
                WebPkiClientVerifier::builder(root_certs.into()).build()?
            })
            .with_single_cert(
                load_cert(&conf.backend_interfaces.tls_cert).await?,
                load_key(&conf.backend_interfaces.tls_key).await?,
            )?;
        server_config.alpn_protocols = vec![b"h2".to_vec(), b"http/1.1".to_vec()];

        axum_server::bind_rustls(
            addr,
            axum_server::tls_rustls::RustlsConfig::from_config(Arc::new(server_config)),
        )
        .serve(app.into_make_service())
        .await?;
    } else {
        axum_server::bind(addr)
            .serve(app.into_make_service())
            .await?;
    }

    Ok(())
}

pub async fn handle_request(b: Bytes) -> Response {
    let b: Vec<u8> = b.into();

    let bp: BasePayload = match serde_json::from_slice(&b) {
        Ok(v) => v,
        Err(e) => {
            return (StatusCode::BAD_REQUEST, e.to_string()).into_response();
        }
    };

    let span = span!(Level::INFO, "request", sender_id = %hex::encode(&bp.sender_id), receiver_id = %hex::encode(&bp.receiver_id), message_type = ?bp.message_type, transaction_id = bp.transaction_id);
    _handle_request(bp, b).instrument(span).await
}

pub async fn _handle_request(bp: BasePayload, b: Vec<u8>) -> Response {
    info!("Request received");

    let sender_client = {
        if bp.sender_id.len() == 8 {
            // JoinEUI.
            let sender_id = match EUI64::from_slice(&bp.sender_id) {
                Ok(v) => v,
                Err(e) => {
                    warn!(error = %e.full(), "Error decoding SenderID as EUI64");
                    let msg = format!("Error decoding SenderID: {}", e);
                    let pl = bp.to_base_payload_result(backend::ResultCode::MalformedRequest, &msg);
                    log_request_response(&bp, &b, &pl).await;
                    return Json(&pl).into_response();
                }
            };

            match joinserver::get(sender_id).await {
                Ok(v) => v,
                Err(_) => {
                    warn!("Unknown SenderID");
                    let msg = format!("Unknown SenderID: {}", sender_id);
                    let pl = bp.to_base_payload_result(backend::ResultCode::UnknownSender, &msg);
                    log_request_response(&bp, &b, &pl).await;
                    return Json(&pl).into_response();
                }
            }
        } else {
            // Unknown size
            warn!("Invalid SenderID length");
            let pl = bp.to_base_payload_result(
                backend::ResultCode::MalformedRequest,
                "Invalid SenderID length",
            );
            log_request_response(&bp, &b, &pl).await;
            return Json(&pl).into_response();
        }
    };

    // Request is an async answer.
    if bp.is_answer() {
        tokio::spawn(async move {
            if let Err(e) = handle_async_ans(&bp, &b).await {
                error!(error = %e.full(), "Handle async answer error");
            }
        });
        return (StatusCode::OK, "").into_response();
    }

    match bp.message_type {
        MessageType::HomeNSReq => handle_home_ns_req(sender_client, bp, &b).await,
        // Unknown message
        _ => (
            StatusCode::INTERNAL_SERVER_ERROR,
            format!("Handler for {:?} is not implemented", bp.message_type),
        )
            .into_response(),
    }
}

fn err_to_response(e: anyhow::Error, bp: &backend::BasePayload) -> backend::BasePayloadResult {
    let msg = format!("{}", e);
    bp.to_base_payload_result(err_to_result_code(e), &msg)
}

fn err_to_result_code(e: anyhow::Error) -> backend::ResultCode {
    if let Some(e) = e.downcast_ref::<StorageError>() {
        return match e {
            StorageError::NotFound(_) => backend::ResultCode::UnknownDevAddr,
            StorageError::InvalidMIC | StorageError::InvalidDevNonce => {
                backend::ResultCode::MICFailed
            }
            _ => backend::ResultCode::Other,
        };
    }
    if let Some(e) = e.downcast_ref::<UplinkError>() {
        return match e {
            UplinkError::RoamingIsNotAllowed => backend::ResultCode::DevRoamingDisallowed,
            _ => backend::ResultCode::Other,
        };
    }
    backend::ResultCode::Other
}

async fn handle_home_ns_req(
    sender_client: Arc<backend::Client>,
    bp: backend::BasePayload,
    b: &[u8],
) -> Response {
    let pl: backend::HomeNSReqPayload = match serde_json::from_slice(b) {
        Ok(v) => v,
        Err(e) => {
            let ans = err_to_response(anyhow::Error::new(e), &bp);
            log_request_response(&bp, b, &ans).await;
            return Json(&ans).into_response();
        }
    };

    if sender_client.is_async() {
        let b = b.to_vec();
        task::spawn(async move {
            let ans = match _handle_home_ns_req(pl).await {
                Ok(v) => v,
                Err(e) => {
                    let msg = e.to_string();
                    backend::HomeNSAnsPayload {
                        base: bp.to_base_payload_result(err_to_result_code(e), &msg),
                        h_net_id: Vec::new(),
                    }
                }
            };

            log_request_response(&bp, &b, &ans).await;

            if let Err(e) = sender_client.home_ns_ans(backend::Role::FNS, &ans).await {
                error!(error = %e.full(), "Send async HomeNSAns error");
            }
        });

        (StatusCode::OK, "").into_response()
    } else {
        match _handle_home_ns_req(pl).await {
            Ok(ans) => {
                log_request_response(&bp, b, &ans).await;
                Json(&ans).into_response()
            }
            Err(e) => {
                let ans = err_to_response(e, &bp);
                log_request_response(&bp, b, &ans).await;
                Json(&ans).into_response()
            }
        }
    }
}

async fn _handle_home_ns_req(pl: backend::HomeNSReqPayload) -> Result<backend::HomeNSAnsPayload> {
    let conf = config::get();

    Ok(backend::HomeNSAnsPayload {
        base: pl
            .base
            .to_base_payload_result(backend::ResultCode::Success, ""),
        h_net_id: conf.network.net_id.to_vec(),
    })
}

async fn handle_async_ans(bp: &BasePayload, b: &[u8]) -> Result<Response> {
    let transaction_id = bp.transaction_id;

    let key = redis_key(format!("backend:async:{}", transaction_id));

    () = redis::pipe()
        .atomic()
        .cmd("XADD")
        .arg(&key)
        .arg("MAXLEN")
        .arg(1_i64)
        .arg("*")
        .arg("pl")
        .arg(b)
        .ignore()
        .cmd("EXPIRE")
        .arg(&key)
        .arg(30_i64)
        .ignore()
        .query_async(&mut get_async_redis_conn().await?)
        .await?;

    Ok((StatusCode::OK, "").into_response())
}

pub async fn get_async_receiver(
    transaction_id: u32,
    timeout: Duration,
) -> Result<oneshot::Receiver<Vec<u8>>> {
    let (tx, rx) = oneshot::channel();

    task::spawn(async move {
        let mut c = match get_async_redis_conn().await {
            Ok(v) => v,
            Err(e) => {
                error!(error = %e, "Get Redis connection error");
                return;
            }
        };
        let key = redis_key(format!("backend:async:{}", transaction_id));

        let srr: StreamReadReply = match redis::cmd("XREAD")
            .arg("BLOCK")
            .arg(timeout.as_millis() as u64)
            .arg("COUNT")
            .arg(1_u64)
            .arg("STREAMS")
            .arg(&key)
            .arg("0")
            .query_async(&mut c)
            .await
        {
            Ok(v) => v,
            Err(e) => {
                error!(error = %e, "Read from Redis Stream error");
                return;
            }
        };

        for stream_key in &srr.keys {
            for stream_id in &stream_key.ids {
                for (k, v) in &stream_id.map {
                    match k.as_ref() {
                        "pl" => {
                            if let redis::Value::BulkString(b) = v {
                                let _ = tx.send(b.to_vec());
                                return;
                            }
                        }
                        _ => {
                            error!(
                                transaction_id = transaction_id,
                                key = %key,
                                "Unexpected key in async stream"
                            );
                        }
                    }
                }
            }
        }
    });

    Ok(rx)
}

async fn log_request_response<T>(bp: &backend::BasePayload, req_body: &[u8], resp: &T)
where
    T: Serialize + BasePayloadResultProvider,
{
    // The incoming request is an async answer.
    // This is already logged by the backend client.
    if bp.is_answer() {
        return;
    }

    let be_req_log = stream_pb::BackendInterfacesRequest {
        sender_id: hex::encode(&bp.sender_id),
        receiver_id: hex::encode(&bp.receiver_id),
        transaction_id: bp.transaction_id,
        message_type: format!("{:?}", bp.message_type),
        request_body: String::from_utf8(req_body.to_vec()).unwrap_or_default(),
        response_body: serde_json::to_string(resp).unwrap_or_default(),
        result_code: format!("{:?}", resp.base_payload().result.result_code),
        time: Some(Utc::now().into()),
        ..Default::default()
    };

    if let Err(e) = stream::backend_interfaces::log_request(be_req_log).await {
        error!(error = %e.full(), "Log Backend Interfaces request error");
    }
}

#[cfg(test)]
pub mod test {
    use super::*;
    use crate::test;

    #[tokio::test]
    async fn test_async_response() {
        let _guard = test::prepare().await;

        let bp = BasePayload {
            transaction_id: 1234,
            ..Default::default()
        };

        let b = vec![1, 2, 3, 4];
        handle_async_ans(&bp, &b).await.unwrap();

        let rx = get_async_receiver(1234, Duration::from_millis(100))
            .await
            .unwrap();

        let rx_b = rx.await.unwrap();
        assert_eq!(b, rx_b);
    }
}
