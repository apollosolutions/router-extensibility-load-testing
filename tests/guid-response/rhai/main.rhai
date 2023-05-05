fn supergraph_service(service){
    service.map_response(|response|{
        for i in 1..=10 {
            response.headers[`GUID${i}`] = uuid_v4();
        };
    });
}