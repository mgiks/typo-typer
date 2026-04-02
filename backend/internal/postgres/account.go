package postgres

import (
	"context"
	"fmt"
)

func (db pgDb) AddAccount(ctx context.Context, username, passhash, salt string) error {
	sql := "INSERT INTO account (username, passhash, salt) VALUES ($1, $2, $3)"
	if _, err := db.pool.Exec(ctx, sql, username, passhash, salt); err != nil {
		return fmt.Errorf("failed to add account: %w", err)
	}
	return nil
}

type account struct {
	id       string
	username string
	email    *string
	passHash string
	salt     string
	wpm      *uint16
}

func (a account) Id() string {
	return a.id
}

func (a account) Username() string {
	return a.username
}

func (a account) Email() *string {
	return a.email
}

func (a account) PassHash() string {
	return a.passHash
}

func (a account) Salt() string {
	return a.salt
}

func (a account) Wpm() *uint16 {
	return a.wpm
}

func (db pgDb) GetAccountByName(ctx context.Context, username string) (account, error) {
	sql := "SELECT id, username, email, passhash, salt, wpm FROM account WHERE username=$1"
	row := db.pool.QueryRow(ctx, sql, username)

	var a account
	if err := row.Scan(&a.id, &a.username, &a.email, &a.passHash, &a.salt, &a.wpm); err != nil {
		return account{}, fmt.Errorf("failed to get account by name: %w", err)
	}

	return a, nil
}
