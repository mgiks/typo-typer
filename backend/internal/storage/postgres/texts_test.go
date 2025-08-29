package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
)

const textFixture = "Test text."

func TestGetRandomText(t *testing.T) {
	db := DB{pool: mockedPool{}}

	text, err := db.GetRandomText(context.Background())
	if err != nil {
		t.Fatal("returned an error")
	}

	if text != textFixture {
		t.Error("should return mocked data")
	}
}

type mockedPool struct{}

func (p mockedPool) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	return mockedRow{}
}

func (p mockedPool) Query(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
	return nil, nil
}

type mockedRow struct{}

func (r mockedRow) Scan(dest ...any) error {
	variable, _ := dest[0].(*string)

	*variable = textFixture

	return nil
}
