package main

import (
	"log"
	"net/http"
	"subgraph-a/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

const port = "8081"

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
