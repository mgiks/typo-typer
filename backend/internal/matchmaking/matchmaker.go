package matchmaking

import (
	"fmt"
	"time"
)

type bucketID int16

type matchMaker struct {
	buckets *bucketMap
	matches *matchMap
}

func NewMatchMaker() *matchMaker {
	return &matchMaker{matches: newMatchMap(), buckets: newBucketMap()}
}

func (mm *matchMaker) Run() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()

	for range t.C {
		for id := range mm.buckets.m {
			mm.matchBucket(id)
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

func (mm *matchMaker) matchBucket(id bucketID) {
	q := mm.buckets.m[id]
	if q == nil {
		return
	}

	for len(q.players) >= 2 {
		p1, _ := q.dequeue()
		p2, _ := q.dequeue()
		mm.matches.createMatch(p1, p2)
	}
}
