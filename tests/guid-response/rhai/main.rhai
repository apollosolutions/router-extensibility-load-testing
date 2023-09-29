fn supergraph_service(service) {
    service.map_response(|response| {
        let guids = [];
        for i in 1..=10 {
            guids += uuid_v4();
        };

        response.headers["GUID"] = guids;
    });
}
