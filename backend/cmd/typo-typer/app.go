package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/logger"
	"github.com/mgiks/typo-typer/internal/text"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/user"
	"github.com/mgiks/typo-typer/internal/validation"
)

type application struct {
	config         config
	wsManager      *WsManager
	textService    text.TextService
	hashingService hashing.HashingService
	userService    user.UserService
	tokenService   token.TokenService
	validator      validation.ValidationService
	logger         logger.LoggerService
}

type config struct {
	port          string
	allowedOrigin string
	db            dbConfig
	ws            wsConfig
	jwt           jwtConfig
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

type jwtConfig struct {
	secret []byte
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
		r.Route("/api", func(r chi.Router) {
			r.Route("/texts", func(r chi.Router) {
				r.Post("/", app.createTextHandler)
				r.Get("/random", app.getRandomTextHandler)
			})

			// TODO: Improve these handlers in the future
			r.Route("/auth", func(r chi.Router) {
				r.Post("/register", app.registerHandler)
				r.Post("/login", app.loginHandler)
			})
		})

		r.Route("/ws", func(r chi.Router) {
			r.Get("/hello", app.helloHandler)
			// r.Route("/matchmaking", func(r chi.Router) {
			// 	r.Get("/pool", app.joinPoolHandler)
			// })
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

	app.logger.Info("server has started", "port", app.config.port)
	return fmt.Errorf("server failed when listening: %w", srv.ListenAndServe())
}
