fn supergraph_service(service){
    service.map_request(|request|{
        log_info("Hello world!");
    });
}