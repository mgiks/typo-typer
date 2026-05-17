package storage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrConflict          = errors.New("resource already exists")
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Store interface {
	Users() UserRepository
	Text() TextRepository
	RefreshToken() RefreshTokenRepository
}

type UserRepository interface {
	Create(context.Context, *User) error
	GetByID(context.Context, int64) (User, error)
	GetByName(context.Context, string) (User, error)
}

type TextRepository interface {
	GetRandom(ctx context.Context) (Text, error)
	Create(context.Context, *Text) error
}

type RefreshTokenRepository interface {
	Create(
		ctx context.Context,
		tokenHash, salt, accountID string,
		expiresAt time.Time,
	) error
}

type store struct {
	user         userStore
	text         textStore
	refreshToken refreshTokenStore
}

func (s store) Users() UserRepository {
	return s.user
}

func (s store) Text() TextRepository {
	return s.text
}

func (s store) RefreshToken() RefreshTokenRepository {
	return s.refreshToken
}

func NewStore(db *pgxpool.Pool) Store {
	return store{
		user:         userStore{db: db},
		text:         textStore{db: db},
		refreshToken: refreshTokenStore{db: db},
	}
}
