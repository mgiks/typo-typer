package server

import (
	"encoding/json"
	"fmt"
	msg "github.com/mgiks/ttyper/internal/message"
)

func (s *server) searchForAMatch(
	p *player,
	message []byte,
) error {
	var unserializedMsg msg.SearchForMatch
	if err := json.Unmarshal(message, &unserializedMsg); err != nil {
		return fmt.Errorf("searchForAMatch: failed to unmarshal message: %w", err)
	}

	p.name = unserializedMsg.Data.PlayerName
	p.id = unserializedMsg.Data.PlayerId

	s.pm.mu.Lock()
	s.pm.searchingPlayers[p.id] = p
	s.pm.mu.Unlock()

	return nil
}
