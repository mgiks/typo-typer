package main

import (
	"context"

	"github.com/coder/websocket"
)

type Client struct {
	id   string
	conn *websocket.Conn
	// Events that are incoming to the client
	incomingEvents chan Event
}

func NewClient(id string, conn *websocket.Conn) Client {
	return Client{
		id:             id,
		conn:           conn,
		incomingEvents: make(chan Event),
	}
}

func (app application) readMessages(c Client) {
	defer func() {
		app.wsManager.removeClient(c.id)
	}()

	c.conn.SetReadLimit(512)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), app.wsManager.readTimeLimit)
		defer cancel()

		var event Event
		if err := readWsJson(ctx, c.conn, &event); err != nil {
			app.logger.Warn("client failed to read json", "err", err)
			return
		}

		if err := app.wsManager.routeEvent(event, c); err != nil {
			app.logger.Warn("manager failed to route event", "err", err)
			internalErrorClose(c.conn)
			return
		}

		cancel()
	}
}

func (app application) writeMessages(c Client) {
	// TODO: Add ping pong handling
	defer func() {
		app.wsManager.removeClient(c.id)
	}()

	for {
		select {
		case event, open := <-c.incomingEvents:
			ctx, cancel := context.WithTimeout(context.Background(), app.wsManager.writeTimeLimit)
			defer cancel()

			if !open {
				normalClosureClose(c.conn)
				return
			}

			if err := writeWsJson(ctx, c.conn, event); err != nil {
				app.logger.Warn("client failed to write json", "err", err)
				internalErrorClose(c.conn)
				return
			}

			cancel()
		}
	}
}
