package matchmaking

import (
	"context"
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
	Names []string `json:"names"`
}

type match struct {
	playerCountTarget uint
	players           matchedPlayers
	msgCh             chan matchMsg
	running           bool
	startedAt         *time.Time
	endsAt            *time.Time
}

func newMatch() *match {
	return &match{
		playerCountTarget: 2,
		players:           make([]*MatchedPlayer, 0, 2),
		msgCh:             make(chan matchMsg, 10),
		running:           false,
		startedAt:         nil,
		endsAt:            nil,
	}
}

func (m *match) run() {
	for msg := range m.msgCh {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		for _, p := range m.players {
			wsjson.Write(ctx, p.conn, msg)
		}
	}
}

func (m *match) start() {
	startTime := time.Now()
	m.startedAt = &startTime

	endTime := startTime.Add(time.Minute * 2)
	m.endsAt = &endTime

	m.msgCh <- matchMsg{Type: "matchStarted"}
}

func (m *match) listenForMsgs(p *MatchedPlayer) {
	for {
		var msg matchMsg
		err := wsjson.Read(context.Background(), p.conn, &msg)
		if err != nil {
			break
		}
		m.msgCh <- msg
	}
}

func (m *match) enter(p *MatchedPlayer) {
	m.players = append(m.players, p)
	go m.listenForMsgs(p)

	if !m.running {
		m.running = true
		go m.run()
	}

	if m.startedAt == nil && len(m.players) >= 2 {
		m.start()
	}

	m.msgCh <- matchMsg{
		Type:    "playerList",
		Payload: playerListPayload{Names: m.players.getNamesSlice()},
	}
}
