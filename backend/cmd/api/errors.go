package main

import (
	"net/http"
)

func (app application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("internal server error", "method", r.Method, "path", r.URL.Path, "err", err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Info("bad request error", "method", r.Method, "path", r.URL.Path, "err", err)
	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app application) unauthorized(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Info("unauthorized error", "method", r.Method, "path", r.URL.Path, "err", err)
	writeJSONError(w, http.StatusUnauthorized, err.Error())
}
