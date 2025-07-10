package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTextHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/texts", nil)
	if err != nil {
		t.Error("failed to create request")
	}

	response := httptest.NewRecorder()

	GETTextHandler(response, request)

	got := response.Body.String()

	if len(got) == 0 {
		t.Error("should respond with non-empty text")
	}
}
