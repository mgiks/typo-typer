package main

import (
	"crypto/rand"
	"net/http"
)

func (app application) joinPoolHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	conn, err := app.acceptWSHandshake(w, r, path)
	if err != nil {
		return
	}

	client := NewClient(rand.Text(), conn, path)
	app.wsManager.addClient(client)

	go app.readMessages(client)
	go app.writeMessages(client)
}
