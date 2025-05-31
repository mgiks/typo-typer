package utils

import (
	"os"
	"testing"
)

func TestFindEnvOr(t *testing.T) {
	t.Run("should call callback if no env found", func(t *testing.T) {
		key := "nonexistentkey"

		unsetEnv(t, key)

		cb := callback{called: false}
		FindEnvOr("nonexistentkey", cb.call)

		if !cb.called {
			t.Errorf("should call callback but doesn't")
		}
	})

	t.Run("should return environmental variable if set", func(t *testing.T) {
		key := "existentkey"
		want := "something"

		setEnv(t, key, want)
		defer unsetEnv(t, key)

		got := FindEnvOr(key, func() {})

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

type callback struct {
	called bool
}

func (c *callback) call() {
	c.called = true
}

func setEnv(t testing.TB, key, value string) {
	t.Helper()
	if os.Setenv(key, value) != nil {
		t.Fatalf("failed to set %q", key)
	}
}

func unsetEnv(t testing.TB, key string) {
	t.Helper()
	if os.Unsetenv(key) != nil {
		t.Fatalf("failed to unset %q", key)
	}
}
