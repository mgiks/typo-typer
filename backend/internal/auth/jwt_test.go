package auth

import (
	"bytes"
	"crypto/ed25519"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)

	data := jwt.MapClaims{
		"user":         "mgik",
		"isADeveloper": true,
	}

	t.Run("should consist of 3 parts separated by dots", func(t *testing.T) {
		token, err := GenerateJWT(seed, data)

		if err != nil {
			t.Errorf("failed to generate jwt: %v", err)
		}

		parts := bytes.Split([]byte(token), []byte("."))

		if len(parts) < 3 {
			t.Errorf("token doesn't consist of 3 parts: %v", token)
		}
	})

	t.Run("should be idempotent", func(t *testing.T) {
		token, err := GenerateJWT(seed, data)
		token2, err2 := GenerateJWT(seed, data)

		if err != nil || err2 != nil {
			t.Errorf("failed to generate jwt: %v", err)
		}

		if token != token2 {
			t.Errorf("the first result %v differs from the second result %v", token, token2)
		}
	})
}
