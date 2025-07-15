package pg

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	vals, err := checkEnvs("POSTGRES_USER", "POSTGRES_PASS", "POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_DB")
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", vals[0], vals[1], vals[2], vals[3], vals[4])

	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return &pgxpool.Pool{}, fmt.Errorf("Connect: failed to connect: %w", err)
	}

	return dbPool, nil
}

func checkEnvs(envs ...string) ([]string, error) {
	vals := make([]string, len(envs))

	for _, key := range envs {
		val := os.Getenv(key)

		if len(val) == 0 {
			return make([]string, 0), errors.New(fmt.Sprint("failed to find environmental variable:", key))
		}

		vals = append(vals, val)
	}

	return vals, nil
}
