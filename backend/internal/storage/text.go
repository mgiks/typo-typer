package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Text struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

type TextStore struct {
	db *pgxpool.Pool
}

func (s TextStore) GetRandom(ctx context.Context) (Text, error) {
	query := `
		SELECT id, content FROM texts 
		ORDER BY RANDOM() 
		LIMIT 1
	`

	var text Text
	if err := s.db.QueryRow(ctx, query).Scan(&text.ID, &text.Content); err != nil {
		return Text{}, err
	}

	return text, nil
}

func (s TextStore) Create(ctx context.Context, text *Text) error {
	query := `
		INSERT INTO texts (content)
		VALUES ($1) RETURNING id
	`
	if err := s.db.QueryRow(ctx, query, text.Content).Scan(&text.ID); err != nil {
		return err
	}
	return nil
}
