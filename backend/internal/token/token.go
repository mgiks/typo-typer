package token

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mgiks/typo-typer/internal/account"
)

type RefreshTokenAdder interface {
	AddRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error
}

type AccountGetter interface {
	GetAccountByName(ctx context.Context, username string) (account.Account, error)
}

type StringHasher interface {
	HashString(str string) (string, string)
}

type TokenService struct {
	privateKey []byte
	a          RefreshTokenAdder
	h          StringHasher
	g          AccountGetter
}

func NewService(privateKey string, a RefreshTokenAdder, h StringHasher, g AccountGetter) (*TokenService, error) {
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("base64 string encoding failed: %w", err)
	}
	return &TokenService{privateKey: key, a: a, h: h, g: g}, nil
}

func (s *TokenService) CreateAccessToken(ctx context.Context, username string) (string, error) {
	account, err := s.g.GetAccountByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find account with such name %s: %w", username, err)
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.Id,
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	str, err := t.SignedString(s.privateKey)
	if err != nil {
		return "", fmt.Errorf("token signing failed: %w", err)
	}
	return str, nil
}

func (s *TokenService) CreateRefreshToken(ctx context.Context, username string) (string, error) {
	token := make([]byte, 32)
	rand.Read(token)
	tokenStr := base64.RawURLEncoding.EncodeToString(token)

	account, err := s.g.GetAccountByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find account with such name %s: %w", username, err)
	}

	tokenHash, salt := s.h.HashString(tokenStr)
	if err := s.a.AddRefreshToken(ctx, tokenHash, salt, account.Id, time.Now().AddDate(0, 0, 30)); err != nil {
		return "", fmt.Errorf("token database insertion failed: %w", err)
	}

	return tokenStr, nil
}
