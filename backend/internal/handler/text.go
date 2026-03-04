package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type getTextResponse struct {
	Text string `json:"text"`
}

type RandomTextGetter interface {
	GetRandomText(ctx context.Context) (string, error)
}

func NewGetTextHandler(g RandomTextGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		text, err := g.GetRandomText(r.Context())
		if err != nil {
			slog.Error("failed to get random text", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		resp := getTextResponse{Text: text}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(resp); err != nil {
			slog.Error("response encoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		if _, err = w.Write(buf.Bytes()); err != nil {
			slog.Error("response write failed", "error", err)
		}
	}
}
