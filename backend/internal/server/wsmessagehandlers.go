package server

import (
	"encoding/json"
	"log"

	"github.com/mgiks/ttyper/internal/dtos"
)

func (s *server) searchForAMatch(
	p *player,
	message []byte,
) {
	var msg dtos.SearchForMatchMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Failed to unmarshal searching player message:", err)
	}

	p.name = msg.Data.PlayerName
	p.id = msg.Data.PlayerId

	s.pm.mu.Lock()
	s.pm.searchingPlayers[p.id] = p
	s.pm.mu.Unlock()
}
