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

func Hash(password string, salt string) string {
	const (
		time    = 1
		memory  = 64 * 1024
		threads = 4
	)

	byteSalt := []byte(salt)

	pepperedPassword := []byte(password + os.Getenv("pepper")) // Pepper is optional

	hash := argon2.IDKey(
		pepperedPassword,
		byteSalt,
		time,
		memory,
		threads,
		32,
	)

	b64Hash := b64.StdEncoding.EncodeToString(hash)
	b64Salt := b64.StdEncoding.EncodeToString(byteSalt)

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

func IsEqualToHash(password string, hash string) (bool, error) {
	hashParts := strings.Split(hash, "$")
	if len(hashParts) < 6 {
		return false, fmt.Errorf("IsEqualToHash: invalid hash")
	}

	salt, err := b64.StdEncoding.DecodeString(hashParts[len(hashParts)-2])
	if err != nil {
		return false, fmt.Errorf("IsEqualToHash: failed to decode salt string: %v", err)
	}

	hashedPass := Hash(password, string(salt))

	return hashedPass == hash, nil
}
