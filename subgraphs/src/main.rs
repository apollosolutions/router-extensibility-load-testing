use std::fs::read_to_string;
use std::sync::Arc;

use serde::{Deserialize, Deserializer};
use axum::{Router, routing::post, Server};
use axum::extract::State;
use axum::response::IntoResponse;

#[tokio::main]
async fn main() {
    let manifest = read_fixtures().expect("could not read fixtures");
    let manifest = Arc::new(manifest);

    let router = Router::new()
        .route("/", post(handler)).with_state(manifest);

    Server::bind(&"0.0.0.0:8082".parse().unwrap()).serve(router.into_make_service()).await.unwrap();
}

async fn handler(State(manifest): State<Arc<Manifest>>, body: String) -> impl IntoResponse {
    let fixture = manifest.fixtures.iter().find(|f| f.req == body);

    let headers = [("content-type", "application/json")];

    let body = if let Some(fixture) = fixture {
        fixture.res.clone()
    } else {
        println!("no fixture for {}", body);
        "{\"errors\":[{\"message\": \"unexpected request\"}]}".to_string()
    };
    (headers, body)
}

#[derive(Debug, Deserialize)]
struct Manifest {
    fixtures: Vec<Fixture>,
}

#[derive(Debug, Deserialize)]
struct Fixture {
    #[serde(deserialize_with = "deserialize_trimmed_string")]
    req: String,
    res: String,
}

fn deserialize_trimmed_string<'de, D>(deserializer: D) -> Result<String, D::Error>
where
    D: Deserializer<'de>,
{
    let s: String = Deserialize::deserialize(deserializer)?;
    Ok(s.trim().to_string())
}

fn read_fixtures() -> Result<Manifest, Box<dyn std::error::Error>> {
    let contents = read_to_string("fixtures.yaml")?;
    let fixtures: Manifest = serde_yaml::from_str(&contents)?;
    Ok(fixtures)
}
