package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRandomTextHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/texts", nil)
	res := httptest.NewRecorder()

	h := NewRandomTextHandler(mockedRandomTextGetter{})
	h(res, req)

	var got TextResponse

	err := json.Unmarshal(res.Body.Bytes(), &got)
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
