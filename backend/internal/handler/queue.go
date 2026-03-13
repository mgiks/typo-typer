package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/mgiks/typo-typer/internal/matchmaking"
)

type queueInitMsg struct {
	Name string  `json:"name"`
	WPM  float32 `json:"wpm"`
}

type queue interface {
	AddPlayer(p *matchmaking.Player)
}

func NewJoinQueueHandler(q queue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{os.Getenv("FRONTEND_URL")},
		})
		if err != nil {
			slog.Error("websocket handshake failed", "error", err)
			return
		}
		defer conn.CloseNow()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var msg queueInitMsg
		err = wsjson.Read(ctx, conn, &msg)
		if err != nil {
			slog.Error("websocket message reading failed", "error", err)
			return
		}

		fmt.Println(msg)

		p := matchmaking.NewPlayer(msg.Name, msg.WPM, conn)
		q.AddPlayer(p)
	}
}
