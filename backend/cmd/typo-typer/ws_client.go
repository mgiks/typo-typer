package main

import (
	"context"
	"fmt"

	"github.com/coder/websocket"
)

type Client struct {
	id      string
	urlPath string
	conn    *websocket.Conn
	// Events that are incoming to the client
	incomingEvents chan Event
}

func NewClient(id string, conn *websocket.Conn, urlPath string) Client {
	return Client{
		id:             id,
		urlPath:        urlPath,
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
			switch err {
			case ErrInvalidPayloadType:
				app.inavalidDataClose(c)
			case ErrConnClosed:
			default:
				app.internalErrorClose(c, fmt.Errorf("failed to read json: %w", err))
			}
			return
		}

		if err := app.wsManager.routeEvent(event, c); err != nil {
			app.internalErrorClose(c, err)
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
				app.normalClosureClose(c)
				return
			}

			if err := writeWsJson(ctx, c.conn, event); err != nil {
				app.internalErrorClose(c, fmt.Errorf("failed to write json: %w", err))
				return
			}

			cancel()
		}
	}
}
