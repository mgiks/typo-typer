package pg

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestConnect(t *testing.T) {
	unsetEnvs(t, pgUser, pgPass, pgHost, pgPort, pgDB)

	_, err := Connect(context.Background())
	if err == nil {
		t.Error("should return an error")
	}
}

func TestGetRandomText(t *testing.T) {
	db := DB{pool: mockedPool{}}

	text, err := db.GetRandomText(context.Background())
	if err != nil {
		t.Fatal("should not return any errors")
	}

	if text != "test text" {
		t.Error("should return mocked data")
	}
}

type (
	mockedPool struct{}
	mockerRow  struct{}
)

func (p mockedPool) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	return mockerRow{}
}

func (r mockerRow) Scan(dest ...any) error {
	variable, _ := dest[0].(*string)

	*variable = "test text"

	return nil
}

func (p mockedPool) Query(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
	return nil, nil
}

func unsetEnvs(t testing.TB, envs ...string) {
	t.Helper()

	for _, env := range envs {
		err := os.Unsetenv(env)
		if err != nil {
			t.Error("failed to unset environmental variable: %w", err)
		}
	}
}
