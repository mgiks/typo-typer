package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Struct struct {
	Client *redis.Client
}

func New(ctx context.Context) *Struct {
	return &Struct{
		Client: redis.NewClient(
			&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
				Protocol: 2,
			},
		),
	}
}
