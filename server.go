package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marklude/go-sovtech/dataloader"
	"github.com/marklude/go-sovtech/graph"
	"github.com/marklude/go-sovtech/graph/generated"
	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		os.Exit(1)
	}

	port := strconv.Itoa(viper.Get("go-sovtech.server.port").(int))
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router := http.NewServeMux()

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, dataloader.Middleware(router)))
}
