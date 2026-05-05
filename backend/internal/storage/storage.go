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
	Account() AccountRepository
	Text() TextRepository
	RefreshToken() RefreshTokenRepository
}

type AccountRepository interface {
	Create(context.Context, *Account) error
	GetByID(context.Context, int64) (Account, error)
	GetByName(context.Context, string) (Account, error)
}

type TextRepository interface {
	GetRandom(ctx context.Context) (Text, error)
}

type RefreshTokenRepository interface {
	Create(
		ctx context.Context,
		tokenHash,
		salt,
		accountId string,
		expiresAt time.Time,
	) error
}

type store struct {
	account      AccountRepository
	text         TextRepository
	refreshToken RefreshTokenRepository
}

func (s store) Account() AccountRepository {
	return s.account
}

func (s store) Text() TextRepository {
	return s.text
}

func (s store) RefreshToken() RefreshTokenRepository {
	return s.refreshToken
}

func NewStore(db *pgxpool.Pool) Store {
	return store{
		account:      AccountStore{db: db},
		text:         TextStore{db: db},
		refreshToken: RefreshTokenStore{db: db},
	}
}
