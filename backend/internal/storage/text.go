package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TextStore struct {
	db *pgxpool.Pool
}

func (ts TextStore) GetRandomText(ctx context.Context) (string, error) {
	sql := "SELECT entry FROM typing_text ORDER BY RANDOM() LIMIT 1"
	row := ts.db.QueryRow(ctx, sql)

	var text string
	if err := row.Scan(&text); err != nil {
		return "", fmt.Errorf("failed to get random text: %w", err)
	}

	return text, nil
}
