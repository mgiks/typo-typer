package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/logger"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	customMiddleware "github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/text"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

type application struct {
	config         config
	textService    text.TextService
	hashingService hashing.HashingService
	accountService account.AccountService
	tokenService   token.TokenService
	matchmaker     *matchmaking.MatchMakingService
	validator      validation.ValidationService
	logger         logger.LoggerService
}

type config struct {
	port string
	db   dbConfig
}

type dbConfig struct {
	url             string
	maxConns        int32
	minIdleConns    int32
	maxConnIdleTime string
}

// TODO: change *chi.Mux to http.Handler once api refactoring is done
func (app application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(customMiddleware.CORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/texts", func(r chi.Router) {
			r.Get("/random", app.getRandomTextHandler)
		})
	})

	return r
}

func (app application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.config.port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("server has started at", "port", app.config.port)
	return fmt.Errorf("server failed when listening: %w", srv.ListenAndServe())
}
