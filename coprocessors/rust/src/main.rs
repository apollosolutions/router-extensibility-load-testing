use axum::{response::IntoResponse, routing::post, Json, Router, Server};
use serde_json::{Map, Value};

#[tokio::main]
async fn main() {
    let router = Router::new()
        .route("/static-subgraph", post(static_subgraph))
        .route("/guid-response", post(guid_response))
        .route("/client-awareness", post(client_awareness));

    Server::bind(&"0.0.0.0:8000".parse().unwrap())
        .serve(router.into_make_service())
        .await
        .unwrap();
}

async fn static_subgraph(Json(mut body): Json<Map<String, Value>>) -> Json<Value> {
    if body.get("stage").is_some_and(|value| {
        value
            .as_str()
            .is_some_and(|stage| stage == "SubgraphRequest")
    }) {
        body.entry("headers")
            .or_insert_with(|| Value::Object(Map::new()))
            .as_object_mut()
            .map(|headers| {
                headers.insert(
                    "source".to_string(),
                    Value::Array(vec!["coprocessor".into()]),
                )
            });
    }
    Json(body.into())
}

async fn guid_response(Json(mut body): Json<Map<String, Value>>) -> Json<Value> {
    if body
        .get("stage")
        .and_then(|value| value.as_str())
        .is_some_and(|stage| stage == "RouterResponse")
    {
        body.entry("headers")
            .or_insert_with(|| Value::Object(Map::new()))
            .as_object_mut()
            .map(|headers| {
                let guids = (0..10)
                    .map(|_| uuid::Uuid::new_v4().to_string())
                    .collect::<Vec<_>>();
                headers.insert("GUID".to_string(), guids.into())
            });
    }
    Json(body.into())
}

async fn client_awareness() -> impl IntoResponse {
    todo!()
}
