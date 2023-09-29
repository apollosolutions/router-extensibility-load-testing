fn subgraph_service(service, subgraph) {
    service.map_request(|request| {
        request.subgraph.headers["source"] = ["rhai"];
    });
}
