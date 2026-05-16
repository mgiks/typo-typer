package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/coder/websocket"
)

func readWsJson(ctx context.Context, conn *websocket.Conn, dest any) error {
	payloadType, payload, err := conn.Read(ctx)
	if payloadType != websocket.MessageText {
		unsupportedDataClose(conn)
		return fmt.Errorf("unsupported data")
	}

	if err != nil {
		abnormalClosureClose(conn)
		return fmt.Errorf("abnormal closure")
	}

	if err := json.Unmarshal(payload, dest); err != nil {
		inavalidDataClose(conn)
		return fmt.Errorf("invalid data")
	}

	return nil
}

func writeWsJson(ctx context.Context, conn *websocket.Conn, data any) error {
	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return conn.Write(ctx, websocket.MessageText, msg)
}
