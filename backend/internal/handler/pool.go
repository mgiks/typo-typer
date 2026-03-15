package handler

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/mgiks/typo-typer/internal/matchmaking"
)

type poolJoinMsg struct {
	Name string `json:"name"`
	WPM  int16  `json:"wpm"`
}

type poolJoiner interface {
	JoinPool(p *matchmaking.Player)
}

func NewJoinPoolHandler(pj poolJoiner) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{os.Getenv("FRONTEND_URL")},
		})
		if err != nil {
			slog.Error("pool joining websocket handshake failed", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var msg poolJoinMsg
		err = wsjson.Read(ctx, conn, &msg)
		if err != nil {
			slog.Error("pool joining websocket message reading failed", "error", err)
			return
		}

		p := matchmaking.NewPlayer(msg.Name, msg.WPM, conn)
		pj.JoinPool(p)
	}
}
