package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	ID       string
	Username string
	Email    *string
	PassHash string
	Salt     string
	WPM      *uint16
}

type AccountStore struct {
	db *pgxpool.Pool
}

// TODO: make this function accept *Account parameter
func (s AccountStore) CreateAccount(ctx context.Context, username, passhash, salt string) error {
	query := `
		INSERT INTO accounts (username, passhash, salt) 
		VALUES ($1, $2, $3)
	`
	if _, err := s.db.Exec(ctx, query, username, passhash, salt); err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

func (s AccountStore) GetAccountByID(ctx context.Context, id int64) (Account, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM accounts 
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
		return Account{}, fmt.Errorf("failed to get account by id: %w", err)
	}

	return a, nil
}

func (s AccountStore) GetAccountByName(ctx context.Context, name string) (Account, error) {
	query := `
		SELECT id, username, email, passhash, salt, wpm FROM accounts
		WHERE name = $1
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
		return Account{}, fmt.Errorf("failed to get account by name: %w", err)
	}

	return a, nil
}
