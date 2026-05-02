package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(ctx context.Context, url string, maxConns, minIdleConns int32, maxConnIdleTime string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse db config: %w", err)
	}

	config.MaxConns = maxConns
	config.MinIdleConns = minIdleConns

	duration, err := time.ParseDuration(maxConnIdleTime)
	if err != nil {
		return nil, fmt.Errorf("failed to parse max connection idle time duration: %w", err)
	}

	config.MaxConnIdleTime = duration

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create new db: %w", err)
	}

	return pool, nil
}
