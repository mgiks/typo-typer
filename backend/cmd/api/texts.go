package main

import (
	"net/http"
)

func (app application) getRandomTextHandler(w http.ResponseWriter, r *http.Request) {
	text, err := app.textService.GetRandomText(r.Context())
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, text); err != nil {
		app.internalServerError(w, r, err)
	}
}
