package middleware

import "net/http"

const allowedOrigin = "http://localhost:5173"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		next.ServeHTTP(w, r)
	})
}
