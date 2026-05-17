package cache

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const keyPrefix = "room:"

type roomStore struct {
	rc *redis.Client
}

func (s roomStore) Get(ctx context.Context, roomID string) {
	s.rc.Get(ctx, keyPrefix+roomID)
}

func (s roomStore) Create(ctx context.Context) (roomID string) {
	roomID = uuid.New().String()
	s.rc.Set(ctx, keyPrefix+roomID, true, time.Minute)
	return
}
