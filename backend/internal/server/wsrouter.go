package server

import (
	"log"
	"sync"
)

type wsMessageHandler func(p *player, message []byte)

type wsMessageRouter struct {
	mhs map[string]wsMessageHandler
	mu  sync.Mutex
}

func newWsMessageRouter() *wsMessageRouter {
	return &wsMessageRouter{mhs: make(map[string]wsMessageHandler)}
}

func (mr *wsMessageRouter) addMessageHandler(
	messageType string,
	mh wsMessageHandler,
) {
	mr.mu.Lock()
	defer mr.mu.Unlock()

	mr.mhs[messageType] = mh
}

func (mr *wsMessageRouter) routeMessage(
	p *player,
	messageType string,
	message []byte,
) {
	messageHandler, exists := mr.mhs[messageType]
	if !exists {
		err := p.conn.CloseNow()
		if err != nil {
			log.Println("Failed to close websocket connection:", err)
		}
	}
	messageHandler(p, message)
}
