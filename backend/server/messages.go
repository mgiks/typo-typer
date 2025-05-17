package server

import (
	"log"

	"github.com/mgiks/ttyper/dtos"
)

func (s *server) createRandomTextMessage() *dtos.RandomTextMessage {
	rtm := dtos.NewRandomTextMessage()
	row := s.db.RandomTextRow()
	if err := row.Scan(
		&rtm.Data.Id,
		&rtm.Data.Text,
		&rtm.Data.Submitter,
		&rtm.Data.Source,
	); err != nil {
		log.Printf("Failed to get random text row: %v\n", err)
		return nil
	}

	return rtm
}

func (s *server) createMatchFoundMessage(
	matchID string,
	playerNames []string,
) *dtos.MatchFoundMessage {
	mfm := dtos.NewMatchFoundMessage()
	row := s.db.RandomText()

	mfm.Data.MatchID = matchID
	mfm.Data.Text = row
	mfm.Data.PlayerNames = playerNames
	return mfm
}
