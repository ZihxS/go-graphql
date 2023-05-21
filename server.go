package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ZihxS/go-graphql/graph"
)

const defaultPort = "1315"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	addr := "localhost:" + port

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
