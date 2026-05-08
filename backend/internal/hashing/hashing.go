package hashing

import (
	"crypto/rand"
	"fmt"
	"slices"

	"golang.org/x/crypto/argon2"
)

type HashingService interface {
	GenerateSalt() []byte
	HashPassword(password string, salt []byte) (hash []byte)
	VerifyPassword(password string, hash, salt []byte) error
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

func (s hashingService) HashPassword(password string, salt []byte) []byte {
	return argon2.IDKey([]byte(password), salt, s.conf.time, s.conf.memory, s.conf.threads, s.conf.keyLength)
}

func (s hashingService) VerifyPassword(password string, hash, salt []byte) error {
	passwordHash := s.HashPassword(password, salt)
	if slices.Compare(passwordHash, hash) != 0 {
		return fmt.Errorf("incorrect password")
	}
	return nil
}

func (s hashingService) GenerateSalt() []byte {
	salt := make([]byte, s.conf.saltLength)
	rand.Read(salt)
	return salt
}
