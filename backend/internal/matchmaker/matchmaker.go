package matchmaker

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mgiks/typo-typer/internal/cache"
)

// How many players are in a room
const roomSize = 2

type MatchMaker interface {
	Run()
	SearchGame(SearchingPlayer)
	ErrorChan() chan error
}

func NewMatchMaker(rr cache.RoomRepository) MatchMaker {
	return &matchMaker{
		players:   []SearchingPlayer{},
		room:      rr,
		errorChan: make(chan error, 10),
	}
}

type matchMaker struct {
	sync.Mutex
	players   []SearchingPlayer
	room      cache.RoomRepository
	errorChan chan error
}

type SearchingPlayer struct {
	Name       string
	RoomIDChan chan string
}

func (mm *matchMaker) Run() {
	go func() {
		ticker := time.NewTicker(time.Second)

		for {
			<-ticker.C

			// TODO: Implement more sophisticated matchmaking logic
			mm.Lock()

			if len(mm.players) < roomSize {
				mm.Unlock()
				continue
			}

			roomID, err := mm.room.Create(context.Background())
			if err != nil {
				mm.errorChan <- fmt.Errorf("failed to create room: %w", err)
				mm.Unlock()
				continue
			}

			matchedPlayers := mm.players[len(mm.players)-roomSize:]

			for _, mp := range matchedPlayers {
				mp.RoomIDChan <- roomID
			}

			mm.players = mm.players[:len(mm.players)-roomSize]

			mm.Unlock()
		}
	}()
}

func (mm *matchMaker) SearchGame(sp SearchingPlayer) {
	mm.Lock()
	mm.players = append(mm.players, sp)
	mm.Unlock()
}

func (mm *matchMaker) ErrorChan() chan error {
	return mm.errorChan
}
