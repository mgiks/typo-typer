package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Database struct {
	*redis.Client
}

func New(ctx context.Context) *Database {
	return &Database{
		redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
				Protocol: 2,
			},
		),
	}
}
