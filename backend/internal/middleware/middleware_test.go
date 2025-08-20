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

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	got := response.Header().Get("Access-Control-Allow-Origin")
	want := "http://localhost:5173"

	if got != want {
		t.Errorf("should set Access-Control-Allow-Origini to %q, got %q", want, got)
	}
}
