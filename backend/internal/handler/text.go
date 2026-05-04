package handler

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/mgiks/typo-typer/internal/text"
)

func NewGetTextHandler(textService text.TextService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text, err := textService.GetRandomText(r.Context())
		if err != nil {
			slog.Error("failed to get random text", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(text); err != nil {
			slog.Error("response encoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		if _, err = w.Write(buf.Bytes()); err != nil {
			slog.Error("response write failed", "error", err)
		}
	}
}
