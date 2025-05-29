package hashing

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt() (string, error) {
	salt := make([]byte, 0)

	max := big.NewInt(256)
	for range 16 {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", fmt.Errorf("GenerateSalt: failed to generate random number: %w", err)
		}

		salt = append(salt, byte(n.Int64()))
	}

	return string(salt), nil
}

func HashAndSalt(password string, salt string) string {
	const (
		time    = 1
		memory  = 64 * 1024
		threads = 4
	)

	pepperedPassword := []byte(password + os.Getenv("pepper")) // Pepper is optional

	hash := argon2.IDKey(
		pepperedPassword,
		[]byte(salt),
		time,
		memory,
		threads,
		32,
	)

	b64Hash := b64.StdEncoding.EncodeToString(hash)
	b64Salt := b64.StdEncoding.EncodeToString([]byte(salt))

	conv := strconv.Itoa
	params := []string{
		"argon2",
		conv(time),
		conv(memory),
		conv(threads),
		b64Salt,
		b64Hash,
	}
	result := strings.Join(params, "$")

	return result
}

func CompareToHash(password string, hash string) bool {
	salt := strings.Split(hash, "$")[4]

	hashedPass := HashAndSalt(password, salt)
	return hashedPass == hash
}
