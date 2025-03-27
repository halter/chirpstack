use anyhow::Result;
use tokio::time::sleep;
use tracing::{error, trace};
use std::time::SystemTime;

use prometheus_client::encoding::EncodeLabelSet;
use prometheus_client::metrics::gauge::Gauge;
use prometheus_client::metrics::family::Family;
use crate::monitoring::prometheus;

use super::data;
use super::multicast as mcast;
use crate::config;
use crate::helpers::errors::PrintFullError;
use crate::storage::{device, multicast};

lazy_static! {
    static ref DUTY_CYCLE_DURATION_GAGUE: Family<(), Gauge> = {
        let counter = Family::<(), Gauge>::default();
        prometheus::register(
            "downlink_duty_cycle_count",
            "Time taken for duty cycle of downlink flow.",
            counter.clone(),
        );
        counter
    };
    static ref DUTY_CYCLE_ITEM_COUNT_GAGUE: Family<(), Gauge> = {
        let counter = Family::<(), Gauge>::default();
        prometheus::register(
            "downlink_duty_cycle_count",
            "Time taken for duty cycle of downlink flow.",
            counter.clone(),
        );
        counter
    };
}

pub async fn class_b_c_scheduler_loop() {
    let conf = config::get();

    loop {
        trace!("Starting class_b_c_scheduler_loop run");

        if let Err(err) = schedule_device_queue_batch(conf.network.scheduler.batch_size).await {
            error!(error = %err, "Scheduling device-queue batch failed");
        } else {
            trace!("class_b_c_scheduler_loop completed successfully");
        }

        sleep(conf.network.scheduler.interval).await;
    }
}

pub async fn multicast_group_queue_scheduler_loop() {
    let conf = config::get();

    loop {
        trace!("Starting multicast-group queue scheduler loop run");

        let cycle_start = SystemTime::now();

        if let Err(err) =
            schedule_multicast_group_queue_batch(conf.network.scheduler.batch_size).await
        {
            error!(error = %err, "Scheduling multicast-group queue batch failed");
        } else {
            trace!("Multicast-group queue scheduler run completed successfully");
        }

        let cycle_duration_millis = SystemTime::now().duration_since(cycle_start).unwrap().as_millis();
        DUTY_CYCLE_DURATION_GAGUE.get_or_create(&()).set(cycle_duration_millis.try_into().unwrap());

        sleep(conf.network.scheduler.interval).await;
    }
}

pub async fn schedule_device_queue_batch(size: usize) -> Result<()> {
    trace!("Getting devices that have schedulable queue-items");
    let devices = device::get_with_class_b_c_queue_items(size).await?;
    trace!(
        device_count = devices.len(),
        "Got this number of devices with schedulable queue-items"
    );

    let mut handles = vec![];

    for dev in devices {
        // Spawn the batch as async tasks.
        let handle = tokio::spawn(async move {
            if let Err(e) = data::Data::handle_schedule_next_queue_item(dev).await {
                error!(error = %e, "Schedule next queue-item for device failed");
            }
        });
        handles.push(handle);
    }

    futures::future::join_all(handles).await;

    Ok(())
}

pub async fn schedule_multicast_group_queue_batch(size: usize) -> Result<()> {
    trace!("Getting schedulable multicast-group queue items");
    let items = multicast::get_schedulable_queue_items(size).await?;
    trace!(
        count = items.len(),
        "Got this number of multicast-group queue items"
    );

    let mut handles = vec![];

    for qi in items {
        let handle = tokio::spawn(async move {
            if let Err(e) = mcast::Multicast::handle_schedule_queue_item(qi).await {
                error!(error = %e.full(), "Schedule multicast-group queue item failed");
            }
        });
        handles.push(handle);
    }

    DUTY_CYCLE_ITEM_COUNT_GAGUE.get_or_create(&()).set(items.len().try_into().unwrap());

    futures::future::join_all(handles).await;
    Ok(())
}
