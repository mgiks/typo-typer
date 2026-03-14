package matchmaking

import "github.com/coder/websocket"

type Player struct {
	name           string
	wpm            int16
	conn           *websocket.Conn
	queueTimeInSec uint64
}

func NewPlayer(name string, wpm int16, conn *websocket.Conn) *Player {
	return &Player{
		name: name,
		wpm:  wpm,
		conn: conn,
	}
}
