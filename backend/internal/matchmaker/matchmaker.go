package matchmaker

import (
	"crypto/rand"
	"sync"
	"time"
)

// How many players are in a room
const roomSize = 2

type MatchMaker interface {
	Run()
	SearchGame(SearchingPlayer)
}

func NewMatchMaker() MatchMaker {
	return &matchMaker{
		players: []SearchingPlayer{},
	}
}

type matchMaker struct {
	sync.Mutex
	players []SearchingPlayer
}

type roomID string

type RoomIDReceiverChan chan roomID

type SearchingPlayer struct {
	Name       string
	RoomIDChan RoomIDReceiverChan
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

			roomID := roomID(rand.Text())
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
