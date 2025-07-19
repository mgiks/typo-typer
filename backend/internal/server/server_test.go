package server

import (
	"context"
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
		t.Fatal("failed to create request")
	}

	getTextHandler := NewGETTextHandler(mockedRandomTextGetter{})

	response := httptest.NewRecorder()

	getTextHandler(response, request)

	var got GETTextResponse

	err = json.Unmarshal(response.Body.Bytes(), &got)
	if err != nil {
		t.Fatal("failed to unmarshal response")
	}

	if len(got.Text) == 0 {
		t.Error("should respond with non-empty text")
	}
}

type mockedRandomTextGetter struct{}

func (g mockedRandomTextGetter) GetRandomText(ctx context.Context) (string, error) {
	return "Test text.", nil
}
