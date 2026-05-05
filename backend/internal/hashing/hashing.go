package hashing

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type HashingService interface {
	HashString(str string) (b64hash string, b64salt string)
	VerifyHash(str, b64hash, b64salt string) error
}

type hashingService struct {
	conf hashingConfig
}

type hashingConfig struct {
	time       uint32
	memory     uint32
	threads    uint8
	keyLength  uint32
	saltLength uint8
}

func NewService() HashingService {
	return hashingService{
		conf: hashingConfig{
			saltLength: 16,
			time:       1,
			memory:     64 * 1024,
			threads:    4,
			keyLength:  32,
		},
	}
}

func (s hashingService) HashString(str string) (b64hash string, b64salt string) {
	salt := generateSalt(s.conf.saltLength)
	hash := argon2.IDKey([]byte(str), salt, s.conf.time, s.conf.memory, s.conf.threads, s.conf.keyLength)
	return b64.StdEncoding.EncodeToString(hash), b64.StdEncoding.EncodeToString(salt)
}

func (s hashingService) VerifyHash(str, b64hash, b64salt string) error {
	salt, err := b64.StdEncoding.DecodeString(b64salt)
	if err != nil {
		return fmt.Errorf("failed to decode salt from b64 form: %w", err)
	}

	strHash := argon2.IDKey([]byte(str), salt, s.conf.time, s.conf.memory, s.conf.threads, s.conf.keyLength)
	b64strHash := b64.StdEncoding.EncodeToString(strHash)
	if b64strHash != b64hash {
		return fmt.Errorf("string's hash is different from provided hash")
	}

	return nil
}

func generateSalt(length uint8) []byte {
	salt := make([]byte, length)
	rand.Read(salt)
	return salt
}
