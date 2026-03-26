package matchmaking

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type msg struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
	sender  string
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

type keyEvent struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

type keyEventBatchPayload struct {
	Events []keyEvent `json:"events"`
}

func newMsg(msgType string, payload any) (*msg, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}
	return &msg{Type: msgType, Payload: p}, nil
}

func (m *msg) setSender(name string) {
	m.sender = name
}

type match struct {
	playerCountTarget uint
	text              string
	msgCh             chan *msg
	running           bool
	startedAt         *time.Time
	endsAt            *time.Time
	psm               playerStateMap
}

func newMatch(text string) *match {
	return &match{
		playerCountTarget: 2,
		text:              text,
		msgCh:             make(chan *msg, 10),
		psm:               make(map[string]*playerState),
	}
}

func (m *match) run() {
	for msg := range m.msgCh {
		switch msg.Type {
		case "keyEventBatch":
			var p keyEventBatchPayload
			if err := json.Unmarshal(msg.Payload, &p); err != nil {
				continue
			}

			ps, ok := m.psm[msg.sender]
			if !ok {
				slog.Error("player state not found", "name", msg.sender)
				return
			}

			ps.applyKeyEvents(p.Events)
		}
	}
}

func (m *match) start() error {
	msg, err := newMsg("matchStarted", nil)
	if err != nil {
		return fmt.Errorf("failed to create match start message: %w", err)
	}

	startTime := time.Now()
	m.startedAt = &startTime

	endTime := startTime.Add(time.Minute * 2)
	m.endsAt = &endTime

	m.msgCh <- msg

	return nil
}

func (m *match) listenForMsgs(p MatchedPlayer) {
	for {
		var msg msg
		err := wsjson.Read(context.Background(), p.conn, &msg)
		if err != nil {
			slog.Error("failed to read json from ws connection", "playerName", p.name, "error", err)
		}

		msg.setSender(p.name)
		m.msgCh <- &msg
	}
}

func (m *match) enter(p MatchedPlayer) error {
	m.psm[p.name] = newPlayerState(p.conn)
	go m.listenForMsgs(p)

	if !m.running {
		m.running = true
		go m.run()
	}

	if m.startedAt == nil && len(m.psm) >= 2 {
		if err := m.start(); err != nil {
			return fmt.Errorf("failed to start match")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	textMsg, err := newMsg("text", textPayload{Text: m.text})
	if err != nil {
		return fmt.Errorf("failed to create text message: %w", err)
	}
	wsjson.Write(ctx, p.conn, textMsg)

	playerListMsg, err := newMsg("playerList", playerListPayload{Names: m.psm.getNamesSlice()})
	if err != nil {
		return fmt.Errorf("failed to create player list message: %w", err)
	}
	m.msgCh <- playerListMsg

	return nil
}

type playerState struct {
	conn           *websocket.Conn
	lastKeyEventId int
	textState      string
	textIndex      int
}

func newPlayerState(conn *websocket.Conn) *playerState {
	return &playerState{conn: conn, lastKeyEventId: -1, textIndex: -1}
}

func (s *playerState) applyKeyEvents(events []keyEvent) {
	var textState strings.Builder
	textState.WriteString(s.textState)

	var lastId int
	for _, e := range events {
		if e.Id < s.lastKeyEventId {
			break
		}
		lastId = e.Id
		textState.WriteString(e.Key)
	}

	s.textState = textState.String()
	s.lastKeyEventId = max(s.lastKeyEventId, lastId)
}

type playerStateMap map[string]*playerState

func (m playerStateMap) getNamesSlice() []string {
	names := make([]string, 0, 2)
	for n := range m {
		names = append(names, n)
	}
	return names
}

type matchMap struct {
	m map[string]*match
}

func newMatchMap() *matchMap {
	return &matchMap{m: make(map[string]*match)}
}

// TODO: make this function variadic(possibly accept more than 2 players )
func (matchMap *matchMap) createMatch(p1, p2 *SearchingPlayer, text string) error {
	id := gonanoid.Must()
	matchMap.m[id] = newMatch(text)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	msg, err := newMsg("matchFound", matchFoundPayload{MatchId: id})
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	wsjson.Write(ctx, p1.conn, msg)
	wsjson.Write(ctx, p2.conn, msg)

	p1.conn.Close(websocket.StatusNormalClosure, "match found")
	p2.conn.Close(websocket.StatusNormalClosure, "match found")

	return nil
}
