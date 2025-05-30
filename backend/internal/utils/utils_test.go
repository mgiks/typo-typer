package utils

import (
	"os"
	"testing"
)

func TestFindEnvOr(t *testing.T) {
	t.Run("should call callback if no env found", func(t *testing.T) {
		key := "nonexistentkey"

		err := os.Unsetenv(key)
		if err != nil {
			t.Errorf("failed to unset %q", key)
		}

		cb := callback{called: false}
		FindEnvOr("nonexistentkey", cb.call)

		if !cb.called {
			t.Errorf("should have called callback but didn't")
		}
	})

	t.Run("should return environmental variable if set", func(t *testing.T) {
		key := "existentkey"
		want := "something"

		err := os.Setenv(key, want)
		if err != nil {
			t.Errorf("failed to set %q", key)
		}

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
