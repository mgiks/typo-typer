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
	pgPass = "POSTGRES_PASS"
	pgHost = "POSTGRES_HOST"
	pgPort = "POSTGRES_PORT"
	pgDB   = "POSTGRES_DB"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	vals, err := checkEnvs(pgUser, pgPass, pgHost, pgPort, pgDB)
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", vals[pgUser], vals[pgPass], vals[pgHost], vals[pgPort], vals[pgDB])

	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	return dbPool, nil
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
