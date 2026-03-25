package database

import (
	"context"
	"fmt"

	"github.com/mgiks/typo-typer/internal/account"
)

func (db *DB) AddAccount(ctx context.Context, username, passhash, salt string) error {
	if _, err := db.pool.Exec(ctx,
		"INSERT INTO account username, passhash, salt VALUES $1, $2, $3",
		username, passhash, salt); err != nil {
		return fmt.Errorf("query add account failed: %w", err)
	}
	return nil
}

func (db *DB) GetAccountByName(ctx context.Context, username string) (*account.Account, error) {
	row := db.pool.QueryRow(ctx, "SELECT id, username, email, passhash, salt, wpm FROM account WHERE username=$1", username)

	var a account.Account
	if err := row.Scan(
		&a.Id,
		&a.Username,
		&a.Email,
		&a.PassHash,
		&a.Salt,
		&a.Wpm,
	); err != nil {
		return nil, fmt.Errorf("query get account failed: %w", err)
	}

	return &a, nil
}
