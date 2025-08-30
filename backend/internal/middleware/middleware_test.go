package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := CORS(next)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	got := res.Header().Get("Access-Control-Allow-Origin")
	want := allowedOrigin

	if got != want {
		t.Errorf("should set Access-Control-Allow-Origini to %q, got %q", want, got)
	}
}
