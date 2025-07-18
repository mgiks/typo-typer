package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	s, err := New(nil)
	if err != nil {
		t.Fatal("failed to create server")
	}

	if s == nil {
		t.Error("should return a non-nil value")
	}
}

func TestGETTextHandler(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/texts", nil)
	if err != nil {
		t.Error("failed to create request")
	}

	response := httptest.NewRecorder()

	GETTextHandler(response, request)

	var got GETTextResponse

	err = json.Unmarshal(response.Body.Bytes(), &got)
	if err != nil {
		t.Error("failed to unmarshal response")
	}

	if len(got.Text) == 0 {
		t.Error("should respond with non-empty text")
	}
}
