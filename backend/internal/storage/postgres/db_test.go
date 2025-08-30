package postgres

import (
	"context"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	unsetEnvs(t, pgUser, pgPass, pgHost, pgPort, pgDB)

	_, err := Connect(context.Background())
	if err == nil {
		t.Error("should return an error")
	}
}

func unsetEnvs(t testing.TB, keys ...string) {
	t.Helper()

	for _, key := range keys {
		err := os.Unsetenv(key)
		if err != nil {
			t.Error("failed to unset environmental variable: %w", err)
		}
	}
}
