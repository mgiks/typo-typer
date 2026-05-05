package account

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/storage"
)

var (
	ErrAccountAlreadyExists = fmt.Errorf("account already exists")
	ErrUsernameEmpty        = fmt.Errorf("non-empty username required")
	ErrPasswordTooShort     = fmt.Errorf("password less than 8 characters in length")
	ErrIncorrectPassword    = fmt.Errorf("incorrect password")
	ErrAccountNotFound      = fmt.Errorf("account not found")
)

type AccountService interface {
	CreateAccount(context.Context, *storage.Account) error
	GetAccountByName(context.Context, string) (storage.Account, error)
	PasswordCorrect(ctx context.Context, username, password string) error
}

type accountService struct {
	account        storage.AccountRepository
	hashingService hashing.HashingService
}

func NewService(repo storage.AccountRepository, hashingService hashing.HashingService) AccountService {
	return accountService{account: repo, hashingService: hashingService}
}

func (s accountService) CreateAccount(ctx context.Context, account *storage.Account) error {
	account.Username = strings.TrimSpace(account.Username)
	if len(account.Username) == 0 {
		return ErrUsernameEmpty
	}
	if len(account.Password) < 8 {
		return ErrPasswordTooShort
	}

	account.PassHash, account.Salt = s.hashingService.HashString(account.Password)

	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	if err := s.account.Create(ctx, account); err != nil {
		switch {
		case errors.Is(err, storage.ErrConflict):
			return ErrAccountAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (s accountService) GetAccountByName(ctx context.Context, name string) (storage.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	a, err := s.account.GetByName(ctx, name)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return storage.Account{}, ErrAccountNotFound
		default:
			return storage.Account{}, err
		}
	}

	return a, nil
}

func (s accountService) PasswordCorrect(ctx context.Context, username, password string) error {
	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	a, err := s.account.GetByName(ctx, username)
	if err != nil {
		return ErrAccountNotFound
	}

	if err := s.hashingService.VerifyHash(password, a.PassHash, a.Salt); err != nil {
		return ErrIncorrectPassword
	}

	return nil
}
