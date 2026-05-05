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
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/storage"
)

type TokenService interface {
	CreateAccessToken(ctx context.Context, username string) (string, error)
	ParseAccessToken(ctx context.Context, tokenString string) (jwt.MapClaims, error)
	CreateRefreshToken(ctx context.Context, username string) (string, error)
}

type tokenService struct {
	privateKey     []byte
	accountService account.AccountService
	hashingService hashing.HashingService
	refreshToken   storage.RefreshTokenRepository
}

func NewService(privateKey string,
	accountService account.AccountService,
	hashingService hashing.HashingService,
) (TokenService, error) {
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return tokenService{}, fmt.Errorf("base64 string encoding failed: %w", err)
	}
	return tokenService{
		privateKey:     key,
		accountService: accountService,
		hashingService: hashingService,
	}, nil
}

func (s tokenService) CreateAccessToken(ctx context.Context, username string) (string, error) {
	account, err := s.accountService.GetAccountByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find account with such name %s: %w", username, err)
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	str, err := t.SignedString(s.privateKey)
	if err != nil {
		return "", fmt.Errorf("token signing failed: %w", err)
	}
	return str, nil
}

func (s tokenService) ParseAccessToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
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

func (s tokenService) CreateRefreshToken(ctx context.Context, username string) (string, error) {
	token := make([]byte, 32)
	rand.Read(token)
	tokenStr := base64.RawURLEncoding.EncodeToString(token)

	account, err := s.accountService.GetAccountByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find account with such name %s: %w", username, err)
	}

	tokenHash, salt := s.hashingService.HashString(tokenStr)
	if err := s.refreshToken.Create(ctx, tokenHash, salt, account.ID, time.Now().AddDate(0, 0, 30)); err != nil {
		return "", fmt.Errorf("token database insertion failed: %w", err)
	}

	return tokenStr, nil
}
