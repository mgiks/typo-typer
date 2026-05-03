package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TextStore struct {
	db *pgxpool.Pool
}

// TODO: make this function return Text
func (s TextStore) GetRandomText(ctx context.Context) (string, error) {
	query := "SELECT entry FROM typing_text ORDER BY RANDOM() LIMIT 1"

	var text string
	if err := s.db.QueryRow(ctx, query).Scan(&text); err != nil {
		return "", fmt.Errorf("failed to get random text: %w", err)
	}

	return text, nil
}
