package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type refreshTokenStore struct {
	db *pgxpool.Pool
}

// TODO: make this function accept *RefreshToken parameter
func (s refreshTokenStore) Create(ctx context.Context, tokenHash, salt, accountID string, expiresAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (token_hash, salt, account_id, expires_at) 
		VALUES ($1, $2, $3, $4)
	`
	if _, err := s.db.Exec(ctx, query, tokenHash, salt, accountID, expiresAt); err != nil {
		return err
	}
	return nil
}
