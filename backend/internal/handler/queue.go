package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/coder/websocket"
	"github.com/mgiks/typo-typer/internal/matchmaking"
)

type queueRequest struct {
	Name string  `json:"name"`
	WPM  float32 `json:"wpm"`
}

type queue interface {
	AddPlayer(p *matchmaking.Player)
}

func NewQueueHandler(q queue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var qr queueRequest
		if err := json.NewDecoder(r.Body).Decode(&qr); err != nil {
			slog.Error("request body decoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{os.Getenv("FRONTEND_URL")},
		})
		defer conn.CloseNow()
		if err != nil {
			slog.Error("websocket handshake failed", "error", err)
			return
		}

		p := matchmaking.NewPlayer(qr.Name, qr.WPM, conn)
		q.AddPlayer(p)
	}
}
