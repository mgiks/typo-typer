package token

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/storage"
	"github.com/mgiks/typo-typer/internal/user"
)

type TokenService interface {
	CreateAccessToken(ctx context.Context, username string) (string, error)
	ParseAccessToken(ctx context.Context, tokenString string) (jwt.MapClaims, error)
	CreateRefreshToken(ctx context.Context, username string) (string, error)
}

type tokenService struct {
	privateKey     []byte
	userService    user.UserService
	hashingService hashing.HashingService
	refreshToken   storage.RefreshTokenRepository
}

func NewService(
	privateKey string,
	userService user.UserService,
	hashingService hashing.HashingService,
	repo storage.RefreshTokenRepository,
) (TokenService, error) {
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return tokenService{}, fmt.Errorf("base64 string encoding failed: %w", err)
	}
	return tokenService{
		privateKey:     key,
		userService:    userService,
		hashingService: hashingService,
		refreshToken:   repo,
	}, nil
}

func (s tokenService) CreateAccessToken(ctx context.Context, username string) (string, error) {
	account, err := s.userService.GetUserByName(ctx, username)
	if err != nil {
		return "", fmt.Errorf("failed to find account with such name %s: %w", username, err)
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": account.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
	})
	str, err := tkn.SignedString(s.privateKey)
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

	ctx, cancel := context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	acc, err := s.userService.GetUserByName(ctx, username)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserNotFound):
			return "", err
		default:
			return "", fmt.Errorf("failed to find account: %w", err)
		}
	}

	salt := s.hashingService.GenerateSalt()
	tokenHash := s.hashingService.HashPassword(tokenStr, salt)

	ctx, cancel = context.WithTimeout(ctx, storage.QueryTimeoutDuration)
	defer cancel()

	if err := s.refreshToken.Create(ctx, string(tokenHash), string(salt), acc.ID, time.Now().AddDate(0, 0, 30)); err != nil {
		return "", fmt.Errorf("failed to create refresh token: %w", err)
	}

	return tokenStr, nil
}
