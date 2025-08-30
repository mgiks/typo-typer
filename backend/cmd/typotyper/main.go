package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/handlers"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/storage/postgres"
)

const port = "8000"

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

	fmt.Println("Listening and serving on port", port)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
