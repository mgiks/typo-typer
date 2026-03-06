package database

import (
	"context"
	"fmt"
	"time"
)

func (db *DB) AddRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error {
	if _, err := db.pool.Exec(ctx,
		"INSERT INTO refresh_token (token_hash, salt, account_id, expires_at) VALUES ($1, $2, $3, $4)",
		tokenHash, salt, accountId, expiresAt); err != nil {
		return fmt.Errorf("query add refresh token failed: %w", err)
	}
	return nil
}
