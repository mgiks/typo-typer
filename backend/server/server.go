package server

import (
	"encoding/json"
	"net/http"
)

type GETTextResponse struct {
	Text string `json:"text"`
}

func GETTextHandler(w http.ResponseWriter, r *http.Request) {
	resp := GETTextResponse{Text: "Some Text"}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(jsonResp); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
