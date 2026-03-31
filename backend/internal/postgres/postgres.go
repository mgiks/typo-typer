package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pgUserEnv = "POSTGRES_USER"
	pgPassEnv = "POSTGRES_PASSWORD"
	pgHostEnv = "POSTGRES_HOST"
	pgPortEnv = "POSTGRES_PORT"
	pgDbEnv   = "POSTGRES_DB"
)

type pgDb struct {
	pool *pgxpool.Pool
}

func Connect(ctx context.Context) (pgDb, error) {
	url, err := buildUrl()
	if err != nil {
		return pgDb{}, fmt.Errorf("failed to build url: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	p, err := pgxpool.New(ctx, url)
	if err != nil {
		return pgDb{}, fmt.Errorf("failed to create new pool connection to pg db: %w", err)
	}

	return pgDb{pool: p}, nil
}

func findEnvs(keys ...string) (map[string]string, error) {
	foundEnvs := make(map[string]string)
	missingEnvs := make([]string, 0)

	for _, key := range keys {
		val := os.Getenv(key)
		if len(val) == 0 {
			missingEnvs = append(missingEnvs, key)
		} else {
			foundEnvs[key] = val
		}
	}

	if len(missingEnvs) > 0 {
		return nil, fmt.Errorf("following environment variables are not set: %v", missingEnvs)
	}

	return foundEnvs, nil
}

func buildUrl() (string, error) {
	envs, err := findEnvs(pgUserEnv, pgPassEnv, pgHostEnv, pgPortEnv, pgDbEnv)
	if err != nil {
		return "", fmt.Errorf("failed to find pg envs: %w", err)
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		envs[pgUserEnv],
		envs[pgPassEnv],
		envs[pgHostEnv],
		envs[pgPortEnv],
		envs[pgDbEnv],
	), nil
}
