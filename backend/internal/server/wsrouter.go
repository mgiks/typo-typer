package server

import (
	"errors"
	"fmt"
)

type wsMessageHandler func(p *player, message []byte) error

type wsMessageRouter struct {
	mhs map[string]wsMessageHandler
}

func newWsMessageRouter() *wsMessageRouter {
	return &wsMessageRouter{mhs: make(map[string]wsMessageHandler)}
}

func (mr *wsMessageRouter) addMessageHandler(
	messageType string,
	mh wsMessageHandler,
) {
	mr.mhs[messageType] = mh
}

func (mr *wsMessageRouter) routeMessage(
	p *player,
	messageType string,
	message []byte,
) error {
	messageHandler, exists := mr.mhs[messageType]
	if !exists {
		return errors.New("routeMessage: failed to find message handler")
	}

	if err := messageHandler(p, message); err != nil {
		return fmt.Errorf("routeMessage: %w", err)
	}

	return nil
}
