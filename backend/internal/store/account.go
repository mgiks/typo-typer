package store

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

func (as AccountStore) CreateAccount(ctx context.Context, username, passhash, salt string) error {
	sql := "INSERT INTO account (username, passhash, salt) VALUES ($1, $2, $3)"
	if _, err := as.db.Exec(ctx, sql, username, passhash, salt); err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

func (as AccountStore) GetAccountByID(ctx context.Context, id int64) (Account, error) {
	sql := "SELECT id, username, email, passhash, salt, wpm FROM account WHERE id=$1"
	row := as.db.QueryRow(ctx, sql, id)

	var a Account
	if err := row.Scan(&a.ID, &a.Username, &a.Email, &a.PassHash, &a.Salt, &a.WPM); err != nil {
		return Account{}, fmt.Errorf("failed to get account by name: %w", err)
	}

	return a, nil
}
