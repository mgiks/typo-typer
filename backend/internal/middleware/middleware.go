package middleware

import "net/http"

const AllowedOrigin = "http://localhost:5173"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", AllowedOrigin)
		next.ServeHTTP(w, r)
	})
}
