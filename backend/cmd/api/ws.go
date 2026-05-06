package main

import (
	"net/http"

	"github.com/coder/websocket"
)

const StatusInavalidMessage = 4000

func (app application) acceptWSHandshake(w http.ResponseWriter, r *http.Request, path string) (*websocket.Conn, error) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: app.config.ws.originPattens,
	})
	if err != nil {
		app.logger.Error(path+": "+"websocket handshake failed", "error", err)
		return nil, err
	}
	return conn, nil
}
