package pg

import (
	"context"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	os.Unsetenv("POSTGRES_DB")
	os.Unsetenv("POSTGRES_PASSWORD")
	os.Unsetenv("POSTGRES_PORT")
	os.Unsetenv("POSTGRES_USER")
	os.Unsetenv("POSTGRES_HOST")

	_, err := Connect(context.Background())
	if err == nil {
		t.Error("should return an error")
	}
}
