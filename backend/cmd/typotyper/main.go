package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/database"
	"github.com/mgiks/typo-typer/internal/handler"
)

const port = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", "error", err)
		return
	}

	ctx := context.Background()

	db, err := database.Connect(ctx)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/texts", handler.NewGetTextHandler(db))

	fmt.Printf("Listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		slog.Error("http listening failed", "error", err)
	}

}
