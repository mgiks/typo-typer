package postgres

import (
	"context"
	"fmt"
)

func (db *DB) GetRandomText(ctx context.Context) (string, error) {
	var text string

	err := db.pool.QueryRow(ctx, `SELECT text FROM "typing_text"`).Scan(&text)
	if err != nil {
		return "", fmt.Errorf("postgres.GetRandomText: failed to query row: %w", err)
	}

	return text, nil
}
