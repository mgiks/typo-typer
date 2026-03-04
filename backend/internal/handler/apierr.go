package handler

import (
	"encoding/json"
	"net/http"
)

type APIErr struct {
	Error string `json:"error"`
}

type APIErrDetailed struct {
	Error   string `json:"error"`
	Details any    `json:"details"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeInternalServerErrorJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(APIErr{
		Error: "internal server error",
	})
}
