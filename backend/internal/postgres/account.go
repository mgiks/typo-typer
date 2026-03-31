package postgres

import (
	"context"
	"fmt"

	"github.com/mgiks/typo-typer/internal/account"
)

func (db pgDb) AddAccount(ctx context.Context, username, passhash, salt string) error {
	sql := "INSERT INTO account (username, passhash, salt) VALUES ($1, $2, $3)"
	if _, err := db.pool.Exec(ctx, sql, username, passhash, salt); err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

func (db pgDb) GetAccountByName(ctx context.Context, username string) (account.Account, error) {
	sql := "SELECT id, username, email, passhash, salt, wpm FROM account WHERE username=$1"
	row := db.pool.QueryRow(ctx, sql, username)

	var a account.Account
	if err := row.Scan(&a.Id, &a.Username, &a.Email, &a.PassHash, &a.Salt, &a.Wpm); err != nil {
		return account.Account{}, fmt.Errorf("failed to get account by name: %w", err)
	}

	return a, nil
}
