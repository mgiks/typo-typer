package matchmaking

import (
	"log/slog"
	"sync"
	"time"
)

type bucketID int16

type matchMaker struct {
	mu      sync.Mutex
	buckets map[bucketID]*queue
}

func NewMatchMaker() *matchMaker {
	return &matchMaker{buckets: make(map[bucketID]*queue)}
}

func (mm *matchMaker) Run() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()

	for range t.C {
		for id := range mm.buckets {
			mm.matchBucket(id)
		}
	}
}

func (mm *matchMaker) Join(p *Player) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	id := bucketID(p.wpm / 10)
	if mm.buckets[id] == nil {
		mm.buckets[id] = &queue{}
	}

	mm.buckets[id].enqueue(p)
}

func (mm *matchMaker) matchBucket(id bucketID) {
	q := mm.buckets[id]
	if q == nil {
		return
	}

	for len(q.players) >= 2 {
		p1, _ := q.dequeue()
		p2, _ := q.dequeue()
		createMatch(p1, p2)
	}
}

func createMatch(p1, p2 *Player) {
	slog.Info("match created", "p1", p1, "p2", p2)
}
