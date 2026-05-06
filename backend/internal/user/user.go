package user

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/storage"
)

var (
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrUsernameEmpty     = fmt.Errorf("non-empty username required")
	ErrPasswordTooShort  = fmt.Errorf("password less than 8 characters in length")
	ErrIncorrectPassword = fmt.Errorf("incorrect password")
	ErrUserNotFound      = fmt.Errorf("user not found")
)

type UserService interface {
	CreateUser(context.Context, *storage.Account) error
	GetUserByName(context.Context, string) (storage.Account, error)
	PasswordCorrect(ctx context.Context, username, password string) error
}

type userService struct {
	user           storage.UserRepository
	hashingService hashing.HashingService
}

func NewService(repo storage.UserRepository, hashingService hashing.HashingService) UserService {
	return userService{user: repo, hashingService: hashingService}
}

func (s userService) CreateUser(ctx context.Context, account *storage.Account) error {
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

	if err := s.user.Create(ctx, account); err != nil {
		switch {
		case errors.Is(err, storage.ErrConflict):
			return ErrUserAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (s userService) GetUserByName(ctx context.Context, name string) (storage.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	a, err := s.user.GetByName(ctx, name)
	if err != nil {
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return storage.Account{}, ErrUserNotFound
		default:
			return storage.Account{}, err
		}
	}

	return a, nil
}

func (s userService) PasswordCorrect(ctx context.Context, username, password string) error {
	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	a, err := s.user.GetByName(ctx, username)
	if err != nil {
		return ErrUserNotFound
	}

	if err := s.hashingService.VerifyHash(password, a.PassHash, a.Salt); err != nil {
		return ErrIncorrectPassword
	}

	return nil
}
