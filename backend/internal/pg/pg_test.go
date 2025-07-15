package pg

import (
	"context"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	os.Unsetenv(pgUser)
	os.Unsetenv(pgPass)
	os.Unsetenv(pgHost)
	os.Unsetenv(pgPort)
	os.Unsetenv(pgDB)

	_, err := Connect(context.Background())
	if err == nil {
		t.Error("should return an error")
	}
}
