package server

import (
	"sync"

	"github.com/coder/websocket"
)

type player struct {
	conn *websocket.Conn
	name string
	id   string
}

type match struct {
	players []*player
}

type playerManager struct {
	searchingPlayers map[string]*player
	mu               sync.Mutex
}

type matchManager struct {
	matches map[string]match
	mu      sync.RWMutex
}

func newPlayerManager() *playerManager {
	return &playerManager{searchingPlayers: make(map[string]*player)}
}

func newMatchManager() *matchManager {
	return &matchManager{matches: make(map[string]match)}
}
