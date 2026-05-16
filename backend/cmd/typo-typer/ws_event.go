package main

import "encoding/json"

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(Event, Client) error

type EventType string

const EventHello EventType = "hello"

func eventHelloHandler(e Event, c Client) error {
	payload, err := json.Marshal("hello")
	if err != nil {
		return err
	}

	event := Event{
		Type:    EventHello,
		Payload: payload,
	}

	c.incomingEvents <- event

	return nil
}
