package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mgiks/typo-typer/internal/texts"
)

type TextResponse struct {
	Text string `json:"text"`
}

func NewRandomTextHandler(g texts.RandomTextGetter) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		t, err := g.GetRandomText(context.TODO())
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		resp := TextResponse{Text: t}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	}
}
