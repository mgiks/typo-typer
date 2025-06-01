package auth

import (
	"crypto/ed25519"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {
	seed := make([]byte, ed25519.SeedSize)
	data := jwt.MapClaims{"user": "mgik"}

	t.Run("should consist of 3 parts separated by dots", func(t *testing.T) {
		token, err := GenerateJWT(seed, data)

		assertNoErrors(t, err)
		assertConsistsOf3Parts(t, token)
	})

	t.Run("should be idempotent", func(t *testing.T) {
		token1, err := GenerateJWT(seed, data)
		token2, err2 := GenerateJWT(seed, data)

		assertNoErrors(t, err, err2)
		assertTokensAreEqual(t, token1, token2)
	})
}

func assertTokensAreEqual(t testing.TB, token1, token2 string) {
	t.Helper()
	if token1 != token2 {
		t.Errorf("%q is not equal to %q", token1, token2)
	}
}

func assertConsistsOf3Parts(t testing.TB, token string) {
	t.Helper()
	if len(strings.Split(token, ".")) != 3 {
		t.Errorf("token doesn't consist of 3 parts: %v", token)
	}
}

func assertNoErrors(t testing.TB, errs ...error) {
	t.Helper()
	for _, err := range errs {
		if err != nil {
			t.Errorf("failed to generate jwt: %v", err)
		}
	}
}
