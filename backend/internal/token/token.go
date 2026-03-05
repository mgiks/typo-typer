package token

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	privateKey []byte
}

func NewService(privateKey string) (*TokenService, error) {
	key, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("base64 string encoding failed: %w", err)
	}
	return &TokenService{privateKey: key}, nil
}

func (s *TokenService) CreateToken(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": jwt.NewNumericDate(time.Now()).Add(10 * time.Minute),
	})
	str, err := t.SignedString(s.privateKey)
	if err != nil {
		return "", fmt.Errorf("token signing failed: %w", err)
	}
	return str, nil
}
