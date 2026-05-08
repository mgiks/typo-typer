package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password password `json:"-"`
	WPM      *uint16  `json:"wpm"`
}

type password struct {
	Text *string
	Hash []byte
	Salt []byte
}

func (p *password) Set(text string, hash []byte, salt []byte) {
	p.Text = &text
	p.Hash = hash
	p.Salt = salt
}

type UserStore struct {
	db *pgxpool.Pool
}

func (s UserStore) Create(ctx context.Context, account *User) error {
	query := `
		INSERT INTO users (username, email, passhash, salt) 
		VALUES ($1, $2, $3, $4) RETURNING id, wpm
	`

	err := s.db.QueryRow(
		ctx,
		query,
		account.Username,
		account.Email,
		account.Password.Hash,
		account.Password.Salt,
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

func (s UserStore) GetByID(ctx context.Context, id int64) (User, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM users 
		WHERE id = $1
	`

	var a User
	err := s.db.QueryRow(ctx, query, id).Scan(
		&a.ID,
		&a.Username,
		&a.Email,
		&a.Password.Hash,
		&a.Password.Salt,
		&a.WPM,
	)
	if err != nil {
		return User{}, err
	}

	return a, nil
}

func (s UserStore) GetByName(ctx context.Context, name string) (User, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM users
		WHERE username = $1
	`

	var a User
	err := s.db.QueryRow(ctx, query, name).Scan(
		&a.ID,
		&a.Username,
		&a.Email,
		&a.Password.Hash,
		&a.Password.Salt,
		&a.WPM,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return User{}, ErrNotFound
		default:
			return User{}, err
		}
	}

	return a, nil
}
