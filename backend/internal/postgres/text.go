package postgres

import (
	"context"
	"fmt"
)

func (db pgDb) GetRandomText(ctx context.Context) (string, error) {
	sql := "SELECT entry FROM typing_text ORDER BY RANDOM() LIMIT 1"
	row := db.pool.QueryRow(ctx, sql)

	var text string
	if err := row.Scan(&text); err != nil {
		return "", fmt.Errorf("failed to get random text: %w", err)
	}

	return text, nil
}
