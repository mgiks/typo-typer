package matchmaking

import (
	"context"
	"fmt"
	"log/slog"
	"time"

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

type matchFoundMsg struct {
	MatchId string `json:"matchId"`
}

func (matchMap *matchMap) createMatch(p1, p2 *SearchingPlayer) {
	id := gonanoid.Must()
	matchMap.m[id] = newMatch()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	msg := matchFoundMsg{MatchId: id}
	wsjson.Write(ctx, p1.conn, msg)
	wsjson.Write(ctx, p2.conn, msg)

	p1.conn.Close(websocket.StatusNormalClosure, "match found")
	p2.conn.Close(websocket.StatusNormalClosure, "match found")

}

type matchMsg struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type playerListPayload struct {
	Name []string `json:"name"`
}

type match struct {
	playerCountTarget uint
	players           matchedPlayers
	msgCh             chan matchMsg
	started           bool
}

func newMatch() *match {
	return &match{
		playerCountTarget: 2,
		players:           make([]*MatchedPlayer, 0, 2),
		msgCh:             make(chan matchMsg, 10),
		started:           false,
	}
}

func (m *match) run() {
	for msg := range m.msgCh {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		for _, p := range m.players {
			wsjson.Write(ctx, p.conn, msg)
			fmt.Println(msg)
		}
	}
}

func (m *match) enter(p *MatchedPlayer) {
	if !m.started {
		m.started = true
		go m.run()
	}
	m.players = append(m.players, p)

	m.msgCh <- matchMsg{
		Type: "playerList",
		Payload: playerListPayload{
			Name: m.players.getNamesSlice(),
		},
	}
}
