package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	PassHash string  `json:"-"`
	Salt     string  `json:"-"`
	WPM      *uint16 `json:"wpm"`
	Password string  `json:"-"`
}

type AccountStore struct {
	db *pgxpool.Pool
}

func (s AccountStore) Create(ctx context.Context, account *Account) error {
	query := `
		INSERT INTO users (username, email, passhash, salt) 
		VALUES ($1, $2, $3, $4) RETURNING id, wpm
	`

	err := s.db.QueryRow(
		ctx,
		query,
		account.Username,
		account.Email,
		account.PassHash,
		account.Salt,
	).Scan(
		&account.ID,
		&account.WPM,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			// 23505 - unique constraint violation error in postgres
			case "23505":
				return ErrConflict
			default:
				return err
			}
		}

	}

	return nil
}

func (s AccountStore) GetByID(ctx context.Context, id int64) (Account, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM users 
		WHERE id = $1
	`

	var a Account
	err := s.db.QueryRow(ctx, query, id).Scan(
		&a.ID,
		&a.Username,
		&a.Email,
		&a.PassHash,
		&a.Salt,
		&a.WPM,
	)
	if err != nil {
		return Account{}, err
	}

	return a, nil
}

func (s AccountStore) GetByName(ctx context.Context, name string) (Account, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM users
		WHERE username = $1
	`

	var a Account
	err := s.db.QueryRow(ctx, query, name).Scan(
		&a.ID,
		&a.Username,
		&a.Email,
		&a.PassHash,
		&a.Salt,
		&a.WPM,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return Account{}, ErrNotFound
		default:
			return Account{}, err
		}
	}

	return a, nil
}
