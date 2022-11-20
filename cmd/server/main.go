package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Boutit/graphql/internal/config"
	"github.com/Boutit/graphql/internal/graph/generated"
	graph "github.com/Boutit/graphql/internal/graph/resolver"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg := config.GetConfig(env)
	
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
			AllowedHeaders: 	[]string{"Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Authorization", "Content-Type"},
			AllowedOrigins:   []string{"http://localhost:8090"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	addr := fmt.Sprintf("%s:%s", cfg.GraphQLConfig.GraphQLHost, cfg.GraphQLConfig.GraphQLHTTPPort)

	s := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	log.Printf("graphql playground available at http://%s", addr)
	log.Printf("graphql service listening on  http://%s/query", addr)
	log.Fatalln(s.ListenAndServe())
}