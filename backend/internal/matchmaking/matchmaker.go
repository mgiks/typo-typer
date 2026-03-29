package matchmaking

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

var ErrMatchNotFound = errors.New("match not found")

type bucketID int16

type randomTextGetter interface {
	GetRandomText(ctx context.Context) (string, error)
}

type matchMaker struct {
	buckets *bucketMap
	matches *matchMap
	tg      randomTextGetter
}

func NewMatchMaker(tg randomTextGetter) *matchMaker {
	return &matchMaker{
		matches: newMatchMap(),
		buckets: newBucketMap(),
		tg:      tg,
	}
}

func (mm *matchMaker) Run() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()

	for range t.C {
		for id := range mm.buckets.m {
			if err := mm.matchBucket(id); err != nil {
				slog.Error("failed to match bucket", "error", err)
			}
		}
	}
}

func (mm *matchMaker) JoinPool(p *SearchingPlayer) {
	mm.buckets.mu.Lock()
	defer mm.buckets.mu.Unlock()

	id := bucketID(p.wpm / 10)
	if mm.buckets.m[id] == nil {
		mm.buckets.m[id] = &queue{}
	}

	mm.buckets.m[id].enqueue(p)
}

func (mm *matchMaker) MatchExists(matchId string) bool {
	_, ok := mm.matches.m[matchId]
	return ok
}

func (mm *matchMaker) PlayerBelongs(matchId, playerId string) bool {
	if ok := mm.MatchExists(matchId); !ok {
		return false
	}
	match := mm.matches.m[matchId]
	return match.playerBelongs(playerId)
}

func (mm *matchMaker) EnterMatch(matchId string, p MatchedPlayer) error {
	if ok := mm.MatchExists(matchId); !ok {
		return ErrMatchNotFound
	}
	match := mm.matches.m[matchId]
	if err := match.enter(p); err != nil {
		return fmt.Errorf("failed to enter match: %w", err)
	}
	return nil
}

func (mm *matchMaker) matchBucket(id bucketID) error {
	q := mm.buckets.m[id]
	if q == nil {
		return fmt.Errorf("bucket with id %d doesn't exist", id)
	}

	text, err := mm.tg.GetRandomText(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get random text: %w", err)
	}

	for len(q.players) >= 2 {
		p1 := q.peek(0)
		p2 := q.peek(1)

		if err := mm.matches.createMatch(text, *p1, *p2); err != nil {
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
