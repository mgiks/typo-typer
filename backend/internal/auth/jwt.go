package auth

import (
	"crypto/ed25519"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(privateKeySeed []byte, data jwt.MapClaims) (string, error) {
	key := ed25519.NewKeyFromSeed(privateKeySeed)
	t := jwt.NewWithClaims(jwt.SigningMethodEdDSA, data)
	return t.SignedString(key)
}
