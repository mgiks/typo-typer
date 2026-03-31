package token

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mgiks/typo-typer/internal/account"
)

type refreshTokenAdder interface {
	AddRefreshToken(ctx context.Context, tokenHash, salt, accountId string, expiresAt time.Time) error
}

type accountGetter interface {
	GetAccountByName(ctx context.Context, username string) (account.Account, error)
}

type stringHasher interface {
	HashString(str string) (string, string)
}

type TokenService struct {
	privateKey []byte
	a          refreshTokenAdder
	h          stringHasher
	g          accountGetter
}

func NewService(privateKey string, a refreshTokenAdder, h stringHasher, g accountGetter) (*TokenService, error) {
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

func (s *TokenService) ParseAccessToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		secret := os.Getenv("JWT_SECRET")
		if len(secret) == 0 {
			return nil, fmt.Errorf("failed to retreive secret")
		}
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse token claims")
	}

	return claims, nil
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
