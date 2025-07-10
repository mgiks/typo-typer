package server

import "net/http"

func GETTextHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Some text")); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}
