package main

import (
	"encoding/json"

	"github.com/mgiks/typo-typer/internal/matchmaker"
)

type (
	EventType string

	Event struct {
		Type    EventType       `json:"type"`
		Payload json.RawMessage `json:"payload"`
	}

	EventHandler func(Event, Client) error
)

const (
	EventJoinPool  EventType = "join_pool"
	EventFoundRoom EventType = "found_room"
)

type (
	EventJoinPoolPayload struct {
		Name string `json:"name"`
		WPM  uint8  `json:"wpm"`
	}

	EventFoundGamePayload struct {
		RoomID string `json:"room_id"`
	}
)

func (app application) eventJoinPoolHandler(e Event, c Client) error {
	var payload EventJoinPoolPayload
	err := json.Unmarshal(e.Payload, &payload)
	if err != nil {
		return err
	}

	roomIDChan := make(chan string)

	app.matchmaker.SearchGame(matchmaker.SearchingPlayer{
		Name:       c.id,
		RoomIDChan: roomIDChan,
	})

	outgoingPayload, err := json.Marshal(EventFoundGamePayload{
		RoomID: string(<-roomIDChan),
	})

	if err != nil {
		return err
	}

	var outgoingEvent Event
	outgoingEvent.Type = EventFoundRoom
	outgoingEvent.Payload = outgoingPayload

	c.incomingEvents <- outgoingEvent
	close(c.incomingEvents)

	return nil
}
