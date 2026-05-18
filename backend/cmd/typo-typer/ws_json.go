package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coder/websocket"
)

var (
	ErrInvalidPayloadType = fmt.Errorf("invalid payload type")
	ErrConnClosed         = fmt.Errorf("connection closed")
)

func readWsJson(ctx context.Context, conn *websocket.Conn, dest any) error {
	payloadType, payload, err := conn.Read(ctx)
	if err != nil {
		if websocket.CloseStatus(err) != -1 {
			return ErrConnClosed
		}
		return fmt.Errorf("error reading from connection: %w", err)
	}

	if payloadType != websocket.MessageText {
		return ErrInvalidPayloadType
	}

	if err := json.Unmarshal(payload, dest); err != nil {
		return fmt.Errorf("json unmarshalling failed: %w", err)
	}

	return nil
}

func writeWsJson(ctx context.Context, conn *websocket.Conn, data any) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("json marshalling failed: %w", err)
	}
	return conn.Write(ctx, websocket.MessageText, msg)
}
