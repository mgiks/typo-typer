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

func (rs RefreshTokenStore) CreateRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error {
	sql := "INSERT INTO refresh_token (token_hash, salt, account_id, expires_at) VALUES ($1, $2, $3, $4)"
	if _, err := rs.db.Exec(ctx, sql, tokenHash, salt, accountId, expiresAt); err != nil {
		return fmt.Errorf("failed to add refresh token: %w", err)
	}
	return nil
}
