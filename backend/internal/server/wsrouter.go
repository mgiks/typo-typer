package server

import (
	"log"
	"sync"
)

type WSMessageHandler func(p *Player, message []byte)

type WSMessageRouter struct {
	MHs map[string]WSMessageHandler
	Mu  sync.Mutex
}

func NewWSMessageRouter() *WSMessageRouter {
	return &WSMessageRouter{MHs: make(map[string]WSMessageHandler)}
}

func (mr *WSMessageRouter) AddMessageHandler(
	messageType string,
	mh WSMessageHandler,
) {
	mr.Mu.Lock()
	defer mr.Mu.Unlock()

	mr.MHs[messageType] = mh
}

func (mr *WSMessageRouter) RouteMessage(
	p *Player,
	messageType string,
	message []byte,
) {
	messageHandler, exists := mr.MHs[messageType]
	if !exists {
		err := p.Conn.CloseNow()
		if err != nil {
			log.Println("Failed to close websocket connection:", err)
		}
	}
	messageHandler(p, message)
}
