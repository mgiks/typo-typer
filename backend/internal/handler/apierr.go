package handler

import (
	"encoding/json"
	"net/http"
)

type apiErr struct {
	Error string `json:"error"`
}

type apiErrDetailed struct {
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
	json.NewEncoder(w).Encode(apiErr{
		Error: "internal server error",
	})
}
