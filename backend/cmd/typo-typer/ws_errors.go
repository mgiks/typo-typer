package main

import (
	"github.com/coder/websocket"
)

const StatusInavalidData = 4000

func unsupportedDataClose(conn *websocket.Conn) {
	conn.Close(websocket.StatusUnsupportedData, "")
}

func inavalidDataClose(conn *websocket.Conn) {
	conn.Close(StatusInavalidData, "invalid data")
}

func abnormalClosureClose(conn *websocket.Conn) {
	conn.Close(websocket.StatusAbnormalClosure, "")
}

func internalErrorClose(conn *websocket.Conn) {
	conn.Close(websocket.StatusInternalError, "internal error")
}
