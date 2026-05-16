package main

import (
	"fmt"
	"net/http"

	"github.com/mgiks/typo-typer/internal/storage"
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

type createTextPayload struct {
	Content string `json:"content"`
}

func (app application) createTextHandler(w http.ResponseWriter, r *http.Request) {
	var payload createTextPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("failed to read json: %w", err))
		return
	}

	text := storage.Text{Content: payload.Content}
	if err := app.textService.CreateText(r.Context(), &text); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to create text: %w", err))
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, text); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to send json response: %w", err))
	}
}
