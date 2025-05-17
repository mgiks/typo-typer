package server

import (
	"sync"

	"github.com/coder/websocket"
)

type Player struct {
	Conn *websocket.Conn
	Name string
	Id   string
}

type Match struct {
	Players []*Player
}

type PlayerManager struct {
	SearchingPlayers map[string]*Player
	Mu               sync.Mutex
}

type MatchManager struct {
	Matches map[string]Match
	Mu      sync.RWMutex
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{SearchingPlayers: make(map[string]*Player)}
}

func NewMatchManager() *MatchManager {
	return &MatchManager{Matches: make(map[string]Match)}
}
