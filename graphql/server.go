package main

import (
	"context"
	"fmt"
	"graphql/database"
	"graphql/graph"
	"graphql/graph/store"
	"graphql/graph/store/mstore"
	"graphql/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.Open()
	if err != nil {
		fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		fmt.Errorf("database is not connected: %w ", err)
	}

	// =========================================================================
	//Initialize Conn layer support
	ms, err := models.NewService(db)
	if err != nil {
		log.Println(err)
	}
	service := mstore.NewService(ms)
	s := store.NewStore(&service)
	err = ms.AutoMigrate()
	if err != nil {
		log.Println(err)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: s}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
