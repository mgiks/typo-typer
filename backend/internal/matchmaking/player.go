package matchmaking

import "github.com/coder/websocket"

type Player struct {
	name           string
	wpm            float32
	conn           *websocket.Conn
	queueTimeInSec uint64
}

func NewPlayer(name string, wpm float32, conn *websocket.Conn) *Player {
	return &Player{
		name: name,
		wpm:  wpm,
		conn: conn,
	}
}
