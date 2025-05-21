package hashing

import (
	"crypto/rand"
	b64 "encoding/base64"
	"log"
	"math/big"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt() (string, error) {
	int := big.NewInt(256)
	ints := make([]byte, 0)

	for range 16 {
		n, err := rand.Int(rand.Reader, int)
		if err != nil {
			log.Println("random number generation failed:", err)
			return "", err
		}
		ints = append(ints, byte(n.Int64()))
	}

	return string(ints), nil
}

func HashAndSalt(password string, salt string) string {
	time := 1
	memory := 64 * 1024
	threads := 4
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(salt),
		uint32(time),
		uint32(memory),
		uint8(threads),
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
	hp := HashAndSalt(password, salt)
	return hp == hash
}
