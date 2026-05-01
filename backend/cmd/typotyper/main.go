package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/db"
	"github.com/mgiks/typo-typer/internal/env"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/middleware"
	"github.com/mgiks/typo-typer/internal/store"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

func main() {
	pg, err := db.Connect(context.Background())
	if err != nil {
		log.Fatalf("database connection failed: %v\n", err)
	}

	defer pg.Close()

	storage := store.NewStore(pg)

	hs := hashing.NewService(hashing.DefaultHashingConfig)
	as := account.NewService(storage.Account(), hs)
	v := validation.NewValidator()
	ts, err := token.NewService(env.GetString("JWT_SECRET", ""), storage, hs)
	if err != nil {
		log.Fatalf("failed to initialize token service: %v\n", err)
		return
	}

	mm := matchmaking.NewMatchMaker(storage.Text())
	go mm.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/texts", handler.NewGetTextHandler(storage.Text()))
	mux.HandleFunc("POST /auth/register", handler.NewRegisterHandler(as, v))
	mux.HandleFunc("POST /auth/login", handler.NewLoginHandler(as, v, ts))
	mux.HandleFunc("GET /matchmaking/pool", handler.NewJoinPoolHandler(mm))
	mux.HandleFunc("GET /matchmaking/match/{matchId}", handler.NewEnterMatchHandler(mm, ts))

	port := env.GetString("PORT", ":8080")
	fmt.Printf("Listening and serving on %s\n", port)

	handler := middleware.CORS(mux)
	if err := http.ListenAndServe(env.GetString("PORT", ":8080"), handler); err != nil {
		log.Fatal("http listening and serving failed: %s", err)
	}
}
