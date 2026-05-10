package matchmaker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/mgiks/typo-typer/internal/text"
)

var ErrMatchNotFound = errors.New("match not found")

type bucketID int16

type MatchMakerService struct {
	pool        *pool
	buckets     *bucketMap
	matches     *matchMap
	textService text.TextService
}

type pool struct {
	mu      sync.Mutex
	players []*websocket.Conn
}

func NewService(textService text.TextService) *MatchMakerService {
	return &MatchMakerService{
		pool:        &pool{players: []*websocket.Conn{}},
		matches:     newMatchMap(),
		buckets:     newBucketMap(),
		textService: textService,
	}
}

func (mm *MatchMakerService) Run() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			<-ticker.C
			mm.pool.mu.Lock()
			if len(mm.pool.players) >= 2 {
				p1 := mm.pool.players[0]
				mm.pool.players = mm.pool.players[1:]
				p2 := mm.pool.players[0]
				mm.pool.players = mm.pool.players[1:]
				p1.Write(context.Background(), websocket.MessageText, []byte(`{"data": { "text": "some text" }}`))
				p2.Write(context.Background(), websocket.MessageText, []byte(`{"data": { "text": "some text" }}`))
			}
			mm.pool.mu.Unlock()
		}
	}()
}

func (mm *MatchMakerService) JoinPool(conn *websocket.Conn) {
	mm.pool.mu.Lock()
	defer mm.pool.mu.Unlock()
	mm.pool.players = append(mm.pool.players, conn)
}

func (mm *MatchMakerService) MatchExists(matchId string) bool {
	_, ok := mm.matches.m[matchId]
	return ok
}

func (mm *MatchMakerService) PlayerBelongs(matchId, playerId string) bool {
	if ok := mm.MatchExists(matchId); !ok {
		return false
	}
	match := mm.matches.m[matchId]
	return match.playerBelongs(playerId)
}

func (mm *MatchMakerService) EnterMatch(matchId string, p MatchedPlayer) error {
	if ok := mm.MatchExists(matchId); !ok {
		return ErrMatchNotFound
	}
	match := mm.matches.m[matchId]
	if err := match.enter(p); err != nil {
		return fmt.Errorf("failed to enter match: %w", err)
	}
	return nil
}

func (mm *MatchMakerService) matchBucket(id bucketID) error {
	q := mm.buckets.m[id]
	if q == nil {
		return fmt.Errorf("bucket with id %d doesn't exist", id)
	}

	text, err := mm.textService.GetRandomText(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get random text: %w", err)
	}

	for len(q.players) >= 2 {
		p1 := q.peek(0)
		p2 := q.peek(1)

		if err := mm.matches.createMatch(text.Content, *p1, *p2); err != nil {
			return fmt.Errorf("failed to create match: %w", err)
		}

		if err := q.dequeue(); err != nil {
			return fmt.Errorf("failed to dequeue player1: %w", err)
		}
		if err := q.dequeue(); err != nil {
			return fmt.Errorf("failed to dequeue player2: %w", err)
		}
	}

	return nil
}
