package database

import (
	"context"
	"fmt"
)

func (db *DB) GetRandomText(ctx context.Context) (string, error) {
	var text string

	if err := db.pool.QueryRow(ctx, "SELECT content from typing_text ORDER BY RANDOM() LIMIT 1").Scan(&text); err != nil {
		return "", fmt.Errorf("query random text failed: %w", err)
	}

	return text, nil
}
