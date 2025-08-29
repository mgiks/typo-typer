package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomTextHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/texts", nil)
	response := httptest.NewRecorder()

	randomTextHandler := NewRandomTextHandler(mockedRandomTextGetter{})
	randomTextHandler(response, request)

	var got TextResponse

	err := json.Unmarshal(response.Body.Bytes(), &got)
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
