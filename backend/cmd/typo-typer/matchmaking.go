package main

import (
	"crypto/rand"
	"net/http"
)

func (app application) joinPoolHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := app.acceptWSHandshake(w, r, r.URL.Path)
	if err != nil {
		return
	}

	client := NewClient(rand.Text(), conn)
	app.wsManager.addClient(client)

	go app.readMessages(client)
	go app.writeMessages(client)
}
