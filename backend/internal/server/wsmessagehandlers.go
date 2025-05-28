package server

import (
	"encoding/json"
	"fmt"
)

func (s *server) searchForAMatch(
	p *player,
	message []byte,
) error {
	var msg searchForMatchMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		return fmt.Errorf("searchForAMatch: failed to unmarshal message: %w", err)
	}

	p.name = msg.Data.PlayerName
	p.id = msg.Data.PlayerId

	s.pm.mu.Lock()
	s.pm.searchingPlayers[p.id] = p
	s.pm.mu.Unlock()

	return nil
}
