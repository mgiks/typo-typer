package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Text struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

type TextStore struct {
	db *pgxpool.Pool
}

func (s TextStore) GetRandomText(ctx context.Context) (Text, error) {
	query := "SELECT id, content FROM texts ORDER BY RANDOM() LIMIT 1"

	var text Text
	if err := s.db.QueryRow(ctx, query).Scan(&text.ID, &text.Content); err != nil {
		return Text{}, fmt.Errorf("failed to get random text: %w", err)
	}

	return text, nil
}
