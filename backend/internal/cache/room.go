package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/redis/go-redis/v9"
)

const (
	keyPrefix = "room:"
	roomTTL   = time.Minute
)

type roomStore struct {
	rc *redis.Client
}

func (s roomStore) Get(ctx context.Context, roomID string) {
	s.rc.Get(ctx, keyPrefix+roomID)
}

func (s roomStore) Create(ctx context.Context) (roomID string, err error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}
	roomID = id.String()
	s.rc.Set(ctx, keyPrefix+roomID, true, roomTTL)
	return
}
