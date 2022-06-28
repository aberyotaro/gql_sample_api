package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aberyotaro/gql_sample_api/graph"
	"github.com/aberyotaro/gql_sample_api/graph/generated"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
