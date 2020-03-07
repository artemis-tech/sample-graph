package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/artemis-tech/sample-graph/graph"
)

func main() {
	port := "3030"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.Handle("/graphql/playground", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{}}),
	))

	log.Printf("connect to http://localhost:%s/graphql/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
