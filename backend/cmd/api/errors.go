package main

import (
	"net/http"
)

func (app application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("internal server error", "method", r.Method, "path", r.URL.Path, "err", err)
	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}
