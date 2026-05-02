package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/db"
	"github.com/mgiks/typo-typer/internal/env"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/storage"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

func main() {
	config := config{
		port: env.GetString("PORT", ":8080"),
		db: dbConfig{
			// url:             "",
			maxConns:        35,
			minIdleConns:    5,
			maxConnIdleTime: time.Minute * 15,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pg, err := db.New(
		ctx,
		config.db.url,
		config.db.maxConns,
		config.db.minIdleConns,
		config.db.maxConnIdleTime,
	)
	if err != nil {
		log.Fatalf("database connection failed: %v\n", err)
	}

	defer pg.Close()

	store := storage.NewStore(pg)

	hashingService := hashing.NewService(hashing.DefaultHashingConfig)
	accountService := account.NewService(store.Account(), hashingService)
	validator := validation.NewService()
	matchmaker := matchmaking.NewService(store.Text())
	tokenService, err := token.NewService(env.GetString("JWT_SECRET", ""), store, hashingService)
	if err != nil {
		log.Fatalf("failed to initialize token service: %v\n", err)
		return
	}

	app := application{
		config:         config,
		hashingService: hashingService,
		accountService: accountService,
		tokenService:   tokenService,
		matchmaker:     matchmaker,
		validator:      validator,
	}

	go app.matchmaker.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/texts", handler.NewGetTextHandler(store.Text()))
	mux.HandleFunc("POST /auth/register", handler.NewRegisterHandler(accountService, validator))
	mux.HandleFunc("POST /auth/login", handler.NewLoginHandler(accountService, validator, tokenService))
	mux.HandleFunc("GET /matchmaking/pool", handler.NewJoinPoolHandler(matchmaker))
	mux.HandleFunc("GET /matchmaking/match/{matchId}", handler.NewEnterMatchHandler(matchmaker, tokenService))

	port := env.GetString("PORT", ":8080")
	fmt.Printf("Listening and serving on %s\n", port)

	handler := middleware.CORS(mux)
	if err := http.ListenAndServe(env.GetString("PORT", ":8080"), handler); err != nil {
		log.Fatalf("http listening and serving failed: %s\n", err)
	}
}
