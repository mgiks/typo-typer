package postgres

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgUser = "POSTGRES_USER"
	pgPass = "POSTGRES_PASSWORD"
	pgHost = "POSTGRES_HOST"
	pgPort = "POSTGRES_PORT"
	pgDB   = "POSTGRES_DB"
)

type Querier interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type DB struct {
	pool Querier
}

func Connect(ctx context.Context) (*DB, error) {
	envs, err := checkEnvs(pgUser, pgPass, pgHost, pgPort, pgDB)
	if err != nil {
		return &DB{}, fmt.Errorf("postgres.Connect: environmental variable failure: %w", err)
	}

	url := buildUrl(envs)

	p, err := pgxpool.New(ctx, url)
	if err != nil {
		return &DB{}, fmt.Errorf("postgres.Connect: failed to create database pool: %w", err)
	}

	db := &DB{pool: p}

	return db, nil
}

func checkEnvs(keys ...string) (map[string]string, error) {
	envs := make(map[string]string)

	for _, key := range keys {
		val := os.Getenv(key)

		if len(val) == 0 {
			return nil, errors.New(fmt.Sprint("failed to find environmental variable:", key))
		}

		envs[key] = val
	}

	return envs, nil
}

func buildUrl(envs map[string]string) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		envs[pgUser],
		envs[pgPass],
		envs[pgHost],
		envs[pgPort],
		envs[pgDB],
	)
}
