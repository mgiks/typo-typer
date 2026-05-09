package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/mgiks/typo-typer/internal/db"
	"github.com/mgiks/typo-typer/internal/env"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/logger"
	"github.com/mgiks/typo-typer/internal/matchmaker"
	"github.com/mgiks/typo-typer/internal/storage"
	"github.com/mgiks/typo-typer/internal/text"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/user"
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
		ws: wsConfig{
			originPattens: env.GetStringSlice("WS_ORIGIN_PATTERNS", []string{"http://localhost:5173"}),
		},
		jwt: jwtConfig{
			secret: env.GetByteSlice("JWT_SECRET", nil),
		},
	}

	logger := logger.NewService(*slog.Default())

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

	logger.Info("database connection established")

	store := storage.NewStore(pg)

	textService := text.NewService(store.Text())
	hashingService := hashing.NewService()
	userService := user.NewService(store.Users(), hashingService)
	validator := validation.NewService()
	matchmaker := matchmaker.NewService(textService)
	tokenService, err := token.NewService(config.jwt.secret, userService, hashingService, store.RefreshToken())
	if err != nil {
		logger.FatalError("failed to initialize token service", "err", err)
		return
	}

	app := application{
		config:         config,
		textService:    textService,
		hashingService: hashingService,
		userService:    userService,
		tokenService:   tokenService,
		matchmaker:     matchmaker,
		validator:      validator,
		logger:         logger,
	}

	go app.matchmaker.Run()

	mux := app.mount()
	app.logger.FatalError("app failed", "error", app.run(mux))
}
