package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Account() AccountRepository
	Text() TextRepository
	RefreshToken() RefreshTokenRepository
}

type AccountRepository interface {
	CreateAccount(ctx context.Context, username, passhash, salt string) error
	GetAccountByID(context.Context, int64) (Account, error)
	GetAccountByName(context.Context, string) (Account, error)
}

type TextRepository interface {
	GetRandomText(ctx context.Context) (string, error)
}

type RefreshTokenRepository interface {
	CreateRefreshToken(
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
