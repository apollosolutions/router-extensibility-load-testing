use axum::{routing::post, Json, Router, Server};
use jsonwebtoken::{decode, DecodingKey, Validation};
use serde::{Deserialize, Serialize};
use serde_json::{json, Map, Value};

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
        .and_then(Value::as_str)
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

async fn client_awareness(Json(mut body): Json<Map<String, Value>>) -> Json<Value> {
    if !body
        .get("stage")
        .and_then(Value::as_str)
        .is_some_and(|stage| stage == "RouterRequest")
    {
        return Json(body.into());
    }

    let Some(token) = body
        .get("headers")
        .and_then(|headers| headers.get("authentication"))
        .and_then(|auth_values| auth_values.get(0))
        .and_then(Value::as_str)
        .and_then(|header| header.strip_prefix("Bearer "))
     else {
        body["control"] = json!({"break": 401});
        return Json(body.into());
    };

    let claims = if let Ok(data) = decode::<Claims>(
        token,
        &DecodingKey::from_secret("apollo".as_ref()),
        &Validation::default(),
    ) {
        data.claims
    } else {
        body["control"] = json!({"break": 401});
        return Json(body.into());
    };
    let client_name = claims
        .client_name
        .unwrap_or_else(|| "coprocessor".to_string());
    let client_version = claims
        .client_version
        .unwrap_or_else(|| "loadtest".to_string());
    if let Some(headers) = body
        .entry("headers")
        .or_insert_with(|| Value::Object(Map::new()))
        .as_object_mut()
    {
        headers.insert(
            "apollographql-client-name".to_string(),
            Value::Array(vec![client_name.into()]),
        );
        headers.insert(
            "apollographql-client-version".to_string(),
            Value::Array(vec![client_version.into()]),
        );
    }
    Json(body.into())
}

#[derive(Debug, Serialize, Deserialize)]
struct Claims {
    exp: usize,
    client_name: Option<String>,
    client_version: Option<String>,
}
