package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/postgres"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
	"log/slog"
	"net/http"
	"os"
)

const (
	defaultPort  = "8080"
	envJWTSecret = "JWT_SECRET"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", "error", err)
		return
	}

	ctx := context.Background()

	pg, err := postgres.Connect(ctx)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return
	}

	hs := hashing.NewService(hashing.DefaultHashingConfig)
	as := account.NewService(pg, hs)
	v := validation.NewValidator()

	secret, ok := os.LookupEnv(envJWTSecret)
	if !ok {
		slog.Error("failed to find environment variable", "name", envJWTSecret)
		return
	}

	ts, err := token.NewService(secret, pg, hs, pg)
	if err != nil {
		slog.Error("token service initialization failed", "error", err)
		return
	}

	mm := matchmaking.NewMatchMaker(pg)
	go mm.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/texts", handler.NewGetTextHandler(pg))
	mux.HandleFunc("POST /auth/register", handler.NewRegisterHandler(as, v))
	mux.HandleFunc("POST /auth/login", handler.NewLoginHandler(as, v, ts))
	mux.HandleFunc("GET /matchmaking/pool", handler.NewJoinPoolHandler(mm))
	mux.HandleFunc("GET /matchmaking/match/{matchId}", handler.NewEnterMatchHandler(mm, ts))

	handler := middleware.CORS(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("Listening and serving on port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		slog.Error("http listening and serving failed", "error", err)
	}

}
