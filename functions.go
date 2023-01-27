package gqlgengofirestore

import (
	"log"
	"net/http"
	"os"

	"bsm.com/gqlgengofirestore/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

const defaultPort = "9090"

func init() {
	// Register an HTTP function with the Functions Framework
	functions.HTTP("MyGQLCustomerFunction", myGqlFunction)
}

// Function myHTTPFunction is an HTTP handler
func myGqlFunction(w http.ResponseWriter, r *http.Request) {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
