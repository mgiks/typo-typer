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
	Create(context.Context, *Account) error
	GetByID(context.Context, int64) (Account, error)
	GetByName(context.Context, string) (Account, error)
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
	user         UserRepository
	text         TextRepository
	refreshToken RefreshTokenRepository
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
		user:         AccountStore{db: db},
		text:         TextStore{db: db},
		refreshToken: RefreshTokenStore{db: db},
	}
}
