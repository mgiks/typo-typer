package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenStore struct {
	db *pgxpool.Pool
}

// TODO: make this function accept *RefreshToken parameter
func (s RefreshTokenStore) CreateRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error {
	query := `
		INSERT INTO refresh_tokens (token_hash, salt, account_id, expires_at) 
		VALUES ($1, $2, $3, $4)
	`
	if _, err := s.db.Exec(ctx, query, tokenHash, salt, accountId, expiresAt); err != nil {
		return fmt.Errorf("failed to add refresh token: %w", err)
	}
	return nil
}
