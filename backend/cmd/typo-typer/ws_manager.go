package main

import (
	"fmt"
	"sync"
	"time"
)

type WsManager struct {
	routes         map[EventType]EventHandler
	clients        map[string]Client
	readTimeLimit  time.Duration
	writeTimeLimit time.Duration
	mu             sync.Mutex
}

func newWsManager(readTimeLimit, writeTimeLimit string) (*WsManager, error) {
	rtl, err := time.ParseDuration(readTimeLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to parse read time limit duration: %w", err)
	}

	wtl, err := time.ParseDuration(writeTimeLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to parse write time limit duration: %w", err)
	}

	return &WsManager{
		routes:         make(map[EventType]EventHandler),
		clients:        make(map[string]Client),
		readTimeLimit:  rtl,
		writeTimeLimit: wtl,
	}, nil
}

func (m *WsManager) addClient(c Client) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.clients[c.id] = c
}

func (m *WsManager) removeClient(clientID string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.clients, clientID)
}

func (m *WsManager) routeEvent(event Event, client Client) error {
	handler, ok := m.routes[event.Type]
	if !ok {
		return fmt.Errorf("no handler for such event type %s", event.Type)
	}

	return handler(event, client)
}
