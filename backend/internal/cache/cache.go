package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheStore interface {
	Room() RoomRepository
}

type RoomRepository interface {
	Get(ctx context.Context, roomID string)
	Create(ctx context.Context) (roomID string)
}

type cacheStore struct {
	room roomStore
}

func NewStore(rc *redis.Client) CacheStore {
	return cacheStore{
		room: roomStore{
			rc: rc,
		},
	}
}

func (s cacheStore) Room() RoomRepository {
	return s.room
}
