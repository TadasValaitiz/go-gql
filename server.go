package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	gen "go-gql-demo/graph/generated"
	resolv "go-gql-demo/graph/resolver"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize resolver
	resolver := &resolv.Resolver{}

	// Create new schema
	c := gen.Config{Resolvers: resolver}
	schema := gen.NewExecutableSchema(c)

	// Create server with custom error presenter
	srv := handler.NewDefaultServer(schema)
	// Add error logging middleware
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// Log the error
		log.Printf("GraphQL error: %v", err)

		// Return detailed error in development
		return &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code":  "INTERNAL_SERVER_ERROR",
				"stack": fmt.Sprintf("%+v", err), // Include stack trace
			},
		}
	})

	// Enable debug mode
	srv.Use(extension.Introspection{})

	// Add CORS middleware
	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Handle the actual request
		srv.ServeHTTP(w, r)
	})

	// Setup HTTP handlers
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Start server
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

const defaultPort = "8080"
