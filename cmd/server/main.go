package main

import (
	"os"

	"github.com/Boutit/graphql/internal/config"
	"github.com/go-chi/chi"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg := config.GetConfig(env)
	
	router := chi.NewRouter()
}