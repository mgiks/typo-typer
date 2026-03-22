package matchmaking

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type msg struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func newMsg(msgType string, payload any) (msg, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return msg{}, fmt.Errorf("failed to marshal payload: %w", err)
	}
	return msg{Type: msgType, Payload: p}, nil
}

type matchFoundPayload struct {
	MatchId string `json:"matchId"`
}

type playerListPayload struct {
	Names []string `json:"names"`
}

type textPayload struct {
	Text string `json:"text"`
}

type typingChunkPayload struct {
	Chunk     string  `json:"chunk"`
	TimeInSec float32 `json:"timeInSec"`
}

type matchMap struct {
	m map[string]*match
}

func newMatchMap() *matchMap {
	return &matchMap{m: make(map[string]*match)}
}

func (matchMap *matchMap) createMatch(p1, p2 *SearchingPlayer, text string) error {
	id := gonanoid.Must()
	matchMap.m[id] = newMatch(text)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	msg, err := newMsg("matchFound", matchFoundPayload{MatchId: id})
	if err != nil {
		return fmt.Errorf("message creation failed: %w", err)
	}

	wsjson.Write(ctx, p1.conn, msg)
	wsjson.Write(ctx, p2.conn, msg)

	p1.conn.Close(websocket.StatusNormalClosure, "match found")
	p2.conn.Close(websocket.StatusNormalClosure, "match found")

	return nil
}

type match struct {
	playerCountTarget uint
	players           matchedPlayers
	text              string
	msgCh             chan msg
	running           bool
	startedAt         *time.Time
	endsAt            *time.Time
}

func newMatch(text string) *match {
	return &match{
		playerCountTarget: 2,
		players:           make([]*MatchedPlayer, 0, 2),
		text:              text,
		msgCh:             make(chan msg, 10),
		running:           false,
		startedAt:         nil,
		endsAt:            nil,
	}
}

func (m *match) run() {
	for msg := range m.msgCh {
		switch msg.Type {
		case "typingChunk":
			var p typingChunkPayload
			json.Unmarshal(msg.Payload, &p)
		}
	}
}

func (m *match) start() {
	startTime := time.Now()
	m.startedAt = &startTime

	endTime := startTime.Add(time.Minute * 2)
	m.endsAt = &endTime

	m.msgCh <- msg{Type: "matchStarted"}
}

func (m *match) listenForMsgs(p *MatchedPlayer) {
	for {
		var msg msg
		wsjson.Read(context.Background(), p.conn, &msg)
		m.msgCh <- msg
	}
}

func (m *match) enter(p *MatchedPlayer) error {
	m.players = append(m.players, p)
	go m.listenForMsgs(p)

	if !m.running {
		m.running = true
		go m.run()
	}

	if m.startedAt == nil && len(m.players) >= 2 {
		m.start()
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	textMsg, err := newMsg("text", textPayload{Text: m.text})
	if err != nil {
		return fmt.Errorf("message creation failed: %w", err)
	}
	wsjson.Write(ctx, p.conn, textMsg)

	playerListMsg, err := newMsg("playerList", playerListPayload{Names: m.players.getNamesSlice()})
	if err != nil {
		return fmt.Errorf("message creation failed: %w", err)
	}
	m.msgCh <- playerListMsg

	return nil
}
