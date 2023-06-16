fn router_service(service) {
    let request_callback = Fn("process_request");
    service.map_request(request_callback);
}

fn process_request(request) {
    let claims = request.context[Router.APOLLO_AUTHENTICATION_JWT_CLAIMS];
    if claims == () {
        throw #{ status: 401 };
    }

    let client_name = "rhai";
    let client_version = "loadtest";
    if "client_name" in claims.keys() {
        client_name = claims["client_name"].to_string();
    }
    if "client_version" in claims.keys() {
        client_version = claims["client_version"].to_string();
    }

    request.headers["apollographql-client-name"] = client_name;
    request.headers["apollographql-client-version"] = client_version;
}
