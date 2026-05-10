package matchmaker

import (
	"github.com/coder/websocket"
)

type SearchingPlayer struct {
	name string
	conn *websocket.Conn
}

func NewSearchingPlayer(name string, wpm uint16, conn *websocket.Conn) *SearchingPlayer {
	return &SearchingPlayer{
		name: name,
		conn: conn,
	}
}

type MatchedPlayer struct {
	name string
	conn *websocket.Conn
}

func NewMatchedPlayer(name string, conn *websocket.Conn) MatchedPlayer {
	return MatchedPlayer{
		name: name,
		conn: conn,
	}
}
