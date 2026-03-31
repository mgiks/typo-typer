package postgres

import (
	"context"
	"fmt"
	"time"
)

func (db pgDb) AddRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error {
	sql := "INSERT INTO refresh_token (token_hash, salt, account_id, expires_at) VALUES ($1, $2, $3, $4)"
	if _, err := db.pool.Exec(ctx, sql, tokenHash, salt, accountId, expiresAt); err != nil {
		return fmt.Errorf("failed to add refresh token: %w", err)
	}
	return nil
}
