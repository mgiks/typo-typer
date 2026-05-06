package main

import "net/http"

func (app application) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", app.config.allowedOrigin)
		next.ServeHTTP(w, r)
	})
}
