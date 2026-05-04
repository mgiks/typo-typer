package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/db"
	"github.com/mgiks/typo-typer/internal/env"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/logger"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/storage"
	"github.com/mgiks/typo-typer/internal/text"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

func main() {
	config := config{
		port: env.GetString("PORT", ":8080"),
		db: dbConfig{
			url:             env.GetString("DB_URL", "postgres://admin:adminpassword@localhost:5433/typo-typer"),
			maxConns:        env.GetInt32("DB_MAX_CONNS", 35),
			minIdleConns:    env.GetInt32("DB_MIN_IDLE_CONNS", 5),
			maxConnIdleTime: env.GetString("DB_MAX_CONN_IDLE_TIME", "15m"),
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

	textService := text.NewService(store.Text())
	hashingService := hashing.NewService(hashing.DefaultHashingConfig)
	accountService := account.NewService(store.Account(), hashingService)
	validator := validation.NewService()
	matchmaker := matchmaking.NewService(store.Text())
	tokenService, err := token.NewService(env.GetString("JWT_SECRET", ""), store, hashingService)
	if err != nil {
		log.Fatalf("failed to initialize token service: %v\n", err)
		return
	}

	logger := logger.NewService(*slog.Default())

	app := application{
		config:         config,
		textService:    textService,
		hashingService: hashingService,
		accountService: accountService,
		tokenService:   tokenService,
		matchmaker:     matchmaker,
		validator:      validator,
		logger:         logger,
	}

	go app.matchmaker.Run()

	mux := app.mount()
	mux.Post("/auth/register", handler.NewRegisterHandler(accountService, validator))
	mux.Post("/auth/login", handler.NewLoginHandler(accountService, validator, tokenService))
	mux.Get("/matchmaking/pool", handler.NewJoinPoolHandler(matchmaker))
	mux.Get("/matchmaking/match/{matchId}", handler.NewEnterMatchHandler(matchmaker, tokenService))

	app.logger.FatalError("app failed", "error", app.run(mux))
}
