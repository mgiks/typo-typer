package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mgiks/typo-typer/internal/pg"
)

type Server struct {
	Pg *pg.DB
}

func New(pgDB *pg.DB) (*Server, error) {
	return &Server{Pg: pgDB}, nil
}

type GETTextResponse struct {
	Text string `json:"text"`
}

func NewGETTextHandler(g pg.RandomTextGetter) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		text, err := g.GetRandomText(context.TODO())
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		resp := GETTextResponse{Text: text}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Println("GETTextHandler: failed to encode json:", err)
			http.Error(w, "", http.StatusInternalServerError)
		}
	}
}
