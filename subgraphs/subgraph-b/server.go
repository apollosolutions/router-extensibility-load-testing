package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"subgraph-b/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%v/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
