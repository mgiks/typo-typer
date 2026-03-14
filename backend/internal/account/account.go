package account

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

var ErrAccountAlreadyExists = errors.New("account already exists")
var ErrUsernameEmpty = errors.New("non-empty username required")
var ErrPasswordTooShort = errors.New("password less than 8 characters in length")
var ErrIncorrectPassword = errors.New("incorrect password")
var ErrAccountNotFound = pgx.ErrNoRows

type Account struct {
	Id       string
	Username string
	Email    *string
	PassHash string
	Salt     string
	Wpm      *uint16
}

type accountRepo interface {
	GetAccountByName(ctx context.Context, username string) (Account, error)
	AddAccount(ctx context.Context, username, passhash, salt string) error
}

type passwordHasher interface {
	HashString(str string) (string, string)
	SameHash(str, hashedStr, salt string) (bool, error)
}

type accountService struct {
	repo   accountRepo
	hasher passwordHasher
}

func NewService(repo accountRepo, hasher passwordHasher) *accountService {
	return &accountService{repo: repo, hasher: hasher}
}

func (s *accountService) CreateAccount(ctx context.Context, username, password string) error {
	username = strings.TrimSpace(username)
	if err := validateUsername(username); err != nil {
		return err
	}

	_, err := s.repo.GetAccountByName(ctx, username)
	if err == nil {
		return ErrAccountAlreadyExists
	}

	if err := validatePassword(password); err != nil {
		return err
	}

	passhash, salt := s.hasher.HashString(password)
	return s.repo.AddAccount(ctx, username, passhash, salt)
}

func (s *accountService) PasswordCorrect(ctx context.Context, username, password string) error {
	if err := validateUsername(username); err != nil {
		return err
	}

	a, err := s.repo.GetAccountByName(ctx, username)
	if err != nil {
		return fmt.Errorf("account retrieving failed: %w", err)
	}

	ok, err := s.hasher.SameHash(password, a.PassHash, a.Salt)
	if err != nil {
		return fmt.Errorf("password correctness check failed: %w", err)
	}
	if !ok {
		return ErrIncorrectPassword
	}

	return nil
}

func validateUsername(username string) error {
	if len(username) == 0 {
		return ErrUsernameEmpty
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	return nil
}
