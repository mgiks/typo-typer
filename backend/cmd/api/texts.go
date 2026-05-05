package main

import (
	"fmt"
	"net/http"
)

func (app application) getRandomTextHandler(w http.ResponseWriter, r *http.Request) {
	text, err := app.textService.GetRandomText(r.Context())
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to get random text: %w", err))
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, text); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to send json response: %w", err))
	}
}
