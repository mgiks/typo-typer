package handler

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type matchEnterMsg struct {
	Name string `json:"name"`
}

type matchEnterer interface {
	EnterMatch(matchId, name string)
}

func NewEnterMatchHandler(me matchEnterer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{os.Getenv("FRONTEND_URL")},
		})
		if err != nil {
			slog.Error("match entering websocket handshake failed", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var msg matchEnterMsg
		err = wsjson.Read(ctx, conn, &msg)
		if err != nil {
			slog.Error("match entering websocket message reading failed", "error", err)
			return
		}

		matchId := r.PathValue("matchId")
		me.EnterMatch(matchId, msg.Name)
	}
}
