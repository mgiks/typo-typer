package matchmaking

import (
	"context"
	"fmt"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type matchMap struct {
	m map[string]*match
}

func newMatchMap() *matchMap {
	return &matchMap{m: make(map[string]*match)}
}

func (matchMap *matchMap) createMatch(p1, p2 *Player) {
	id := gonanoid.Must()
	matchMap.m[id] = newMatch()

	ctx := context.Background()
	wsjson.Write(ctx, p1.conn, id)
	wsjson.Write(ctx, p2.conn, id)

	p1.conn.Close(websocket.StatusNormalClosure, "match found")
	p2.conn.Close(websocket.StatusNormalClosure, "match found")
}

type match struct {
	playerCountTarget uint
	players           []string
	msgCh             chan any
}

func newMatch() *match {
	return &match{
		playerCountTarget: 2,
		players:           make([]string, 0, 2),
		msgCh:             make(chan any),
	}
}

func (m *match) enter(name string) {
	m.players = append(m.players, name)
	m.playerCountTarget++
	m.msgCh <- "player entered"
	if len(m.players) == 1 {
		go m.run()
	}
}

func (m *match) run() {
	for msg := range m.msgCh {
		fmt.Println(msg)
	}
}
