package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tonyghita/gqlgen-example/graph"
	"github.com/tonyghita/gqlgen-example/swapi"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	app := graph.Application{
		Client: swapi.NewClient(http.DefaultClient),
	}

	http.Handle("/", handler.Playground("StarWars API", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
