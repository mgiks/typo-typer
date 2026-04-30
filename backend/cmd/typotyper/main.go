package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/postgres"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

const (
	defaultPort  = "8080"
	envJWTSecret = "JWT_SECRET"
)

func main() {
	jwtSecret, ok := os.LookupEnv(envJWTSecret)
	if !ok {
		log.Fatalf("failed to find environment variable %v\n", envJWTSecret)
	}

	ctx := context.Background()
	pg, err := postgres.Connect(ctx)
	if err != nil {
		log.Fatalf("database connection failed: %v\n", err)
	}

	hs := hashing.NewService(hashing.DefaultHashingConfig)
	as := account.NewService(pg, hs)
	v := validation.NewValidator()
	ts, err := token.NewService(jwtSecret, pg, hs, pg)
	if err != nil {
		log.Fatalf("failed to initialize token service: %v\n", err)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Printf("Listening and serving on port %s\n", port)

	handler := middleware.CORS(mux)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("http listening and serving failed: %s", err)
	}
}
