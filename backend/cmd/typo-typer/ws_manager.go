package main

import (
	"fmt"
	"sync"
)

type WsManager struct {
	routes  map[EventType]EventHandler
	clients map[string]Client
	mu      sync.Mutex
}

func newWsManager() *WsManager {
	return &WsManager{
		routes:  make(map[EventType]EventHandler),
		clients: make(map[string]Client),
	}
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

func (m *WsManager) mount() {
	m.routes[EventHello] = eventHelloHandler
}
