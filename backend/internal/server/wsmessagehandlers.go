package server

import (
	"encoding/json"
	"log"

	"github.com/mgiks/ttyper/internal/dtos"
)

func (s *server) SearchForAMatch(
	p *Player,
	message []byte,
) {
	var msg dtos.SearchForMatchMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Println("Failed to unmarshal searching player message:", err)
	}

	p.Name = msg.Data.PlayerName
	p.Id = msg.Data.PlayerId

	s.pm.Mu.Lock()
	s.pm.SearchingPlayers[p.Id] = p
	s.pm.Mu.Unlock()
}
