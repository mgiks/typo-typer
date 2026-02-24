package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgUser = "POSTGRES_USER"
	pgPass = "POSTGRES_PASSWORD"
	pgHost = "POSTGRES_HOST"
	pgPort = "POSTGRES_PORT"
	pgDB   = "POSTGRES_DB"
)

type DB struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context) (*DB, error) {
	envs, err := findEnvs(pgUser, pgPass, pgHost, pgPort, pgDB)
	if err != nil {
		return nil, fmt.Errorf("database configuration error: %w", err)
	}

	dbpool, err := pgxpool.New(context.Background(), buildUrl(envs))
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return &DB{pool: dbpool}, nil
}

func findEnvs(keys ...string) (map[string]string, error) {
	foundEnvs := make(map[string]string)
	missingEnvs := make([]string, 0)

	for _, keyName := range keys {
		val := os.Getenv(keyName)
		if len(val) == 0 {
			missingEnvs = append(missingEnvs, keyName)
		} else {
			foundEnvs[keyName] = val
		}
	}

	if len(missingEnvs) == 0 {
		return foundEnvs, nil
	}

	return nil, fmt.Errorf("following environment variables are not set: %v", missingEnvs)
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
