package server

import (
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

func GETTextHandler(w http.ResponseWriter, r *http.Request) {
	resp := GETTextResponse{Text: "Some Text"}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("GETTextHandler: failed to encode json:", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
