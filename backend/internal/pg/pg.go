package pg

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
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
	envs, err := checkEnvs(pgUser, pgPass, pgHost, pgPort, pgDB)
	if err != nil {
		return &DB{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", envs[pgUser], envs[pgPass], envs[pgHost], envs[pgPort], envs[pgDB])

	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return &DB{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	db := &DB{
		pool: dbPool,
	}

	return db, nil
}

func checkEnvs(envs ...string) (map[string]string, error) {
	keyValPairs := make(map[string]string)

	for _, key := range envs {
		val := os.Getenv(key)

		if len(val) == 0 {
			return make(map[string]string), errors.New(fmt.Sprint("failed to find environmental variable:", key))
		}

		keyValPairs[key] = val
	}

	return keyValPairs, nil
}
