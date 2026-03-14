package matchmaking

import "fmt"

type queue struct {
	players []*Player
}

func (q *queue) enqueue(p *Player) {
	q.players = append([]*Player{p}, q.players...)
}

func (q *queue) dequeue() (*Player, error) {
	if len(q.players) == 0 {
		return nil, fmt.Errorf("unable to dequeue from empty queue")
	}
	p := q.players[len(q.players)-1]
	q.players = q.players[:len(q.players)-1]
	return p, nil
}
