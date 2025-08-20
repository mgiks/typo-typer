package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/pg"
	"github.com/mgiks/typo-typer/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	pgDB, err := pg.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	s, err := server.New(pgDB)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /texts", server.NewGETTextHandler(s.Pg))

	handler := middleware.CORS(mux)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
