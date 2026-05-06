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
	"github.com/mgiks/typo-typer/internal/matchmaker"
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
	matchmaker     *matchmaker.MatchMakerService
	validator      validation.ValidationService
	logger         logger.LoggerService
}

type config struct {
	port          string
	allowedOrigin string
	db            dbConfig
	ws            wsConfig
}

type dbConfig struct {
	url             string
	maxConns        int32
	minIdleConns    int32
	maxConnIdleTime string
}

type wsConfig struct {
	originPattens []string
}

func (app application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(app.CORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/texts", func(r chi.Router) {
			r.Get("/random", app.getRandomTextHandler)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", app.registerHandler)
			r.Post("/login", app.loginHandler)
		})

		r.Route("/matchmaking", func(r chi.Router) {
			r.Get("/pool", app.joinPoolHandler)
			r.Get("/match/{matchID}", app.enterMatchHandler)
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
