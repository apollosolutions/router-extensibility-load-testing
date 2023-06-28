use hyper::body::Bytes;
use serde::{Deserialize, Deserializer};
use std::sync::Arc;
use tokio::{fs::File, io::AsyncReadExt};
use warp::{http::Response, Filter};

#[tokio::main]
async fn main() {
    let manifest = read_fixtures().await.expect("could not read fixtures");
    let manifest = Arc::new(manifest);

    let fixed = warp::post()
        .and(warp::path::end())
        .and(warp::body::bytes())
        .map(move |body: Bytes| {
            let req = String::from_utf8(body.to_vec()).expect("invalid utf8");
            // println!("request: {}", req);

            let pair = manifest.fixtures.iter().find(|f| f.req == req);

            match pair {
                Some(f) => {
                    // println!("found response");
                    Response::builder()
                        .header("content-type", "application/json")
                        .body(f.res.clone())
                }
                None => {
                    // println!("no fixture for request: {}", req);
                    Response::builder()
                        .header("content-type", "application/json")
                        .body(String::from(
                            "{\"errors\":[{\"message\": \"unexpected request\"}]}",
                        ))
                }
            }
        });

    warp::serve(fixed).run(([0, 0, 0, 0], 8082)).await;
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

async fn read_fixtures() -> Result<Manifest, Box<dyn std::error::Error>> {
    let mut file = File::open("fixtures.yaml").await?;
    let mut contents = String::new();
    file.read_to_string(&mut contents).await?;

    let fixtures: Manifest = serde_yaml::from_str(&contents)?;

    Ok(fixtures)
}
