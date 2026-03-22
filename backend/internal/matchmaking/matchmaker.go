package matchmaking

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

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

func (mm *matchMaker) EnterMatch(matchId string, p *MatchedPlayer) error {
	match, ok := mm.matches.m[matchId]
	if !ok {
		return fmt.Errorf("match with id %s not found", matchId)
	}
	match.enter(p)
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
		p1, _ := q.dequeue()
		p2, _ := q.dequeue()
		mm.matches.createMatch(p1, p2, text)
	}

	return nil
}
