use anyhow::Result;

pub mod joinserver;
pub mod keywrap;

pub async fn setup() -> Result<()> {
    joinserver::setup().await?;

    Ok(())
}
