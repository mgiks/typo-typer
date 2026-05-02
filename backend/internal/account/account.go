package account

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mgiks/typo-typer/internal/storage"
)

var ErrAccountAlreadyExists = fmt.Errorf("account already exists")
var ErrUsernameEmpty = fmt.Errorf("non-empty username required")
var ErrPasswordTooShort = fmt.Errorf("password less than 8 characters in length")
var ErrIncorrectPassword = fmt.Errorf("incorrect password")
var ErrAccountNotFound = fmt.Errorf("account not found")

type Account struct {
	Id       string
	Username string
	Email    *string
	PassHash string
	Salt     string
	Wpm      *uint16
}

type passwordHasher interface {
	HashString(str string) (b64hash string, b64salt string)
	VerifyHash(str, b64hash, b64salt string) error
}

type accountService struct {
	repo storage.AccountRepository
	ph   passwordHasher
}

func NewService(repo storage.AccountRepository, ph passwordHasher) *accountService {
	return AccountService{repo: repo, ph: ph}
}

func (s *accountService) CreateAccount(ctx context.Context, username, password string) error {
	username = strings.TrimSpace(username)
	if len(username) == 0 {
		return ErrUsernameEmpty
	}
	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if _, err := s.repo.GetAccountByName(ctx, username); err == nil {
		return ErrAccountAlreadyExists
	}

	passhash, salt := s.ph.HashString(password)
	return s.repo.CreateAccount(ctx, username, passhash, salt)
}

func (s *accountService) PasswordCorrect(ctx context.Context, username, password string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	a, err := s.repo.GetAccountByName(ctx, username)
	if err != nil {
		return ErrAccountNotFound
	}

	if err := s.ph.VerifyHash(password, a.PassHash, a.Salt); err != nil {
		return ErrIncorrectPassword
	}

	return nil
}
