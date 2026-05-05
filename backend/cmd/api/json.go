package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app application) jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}
	if err := writeJSON(w, status, envelope{Data: data}); err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return writeJSON(w, status, envelope{
		Error: message,
	})
}

func readJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dest)
}
