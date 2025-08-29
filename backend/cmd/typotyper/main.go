package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/handlers"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/storage/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB, err := postgres.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /texts", handlers.NewRandomTextHandler(DB))

	handler := middleware.CORS(mux)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
