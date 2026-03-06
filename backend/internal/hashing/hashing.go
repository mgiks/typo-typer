package hashing

import (
	"crypto/rand"
	b64 "encoding/base64"

	"golang.org/x/crypto/argon2"
)

type HashingConfig struct {
	time       uint32
	memory     uint32
	threads    uint8
	keyLength  uint32
	saltLength uint8
}

var DefaultHashingConfig = HashingConfig{
	saltLength: 16,
	time:       1,
	memory:     64 * 1024,
	threads:    4,
	keyLength:  32,
}

type HashingService struct {
	conf HashingConfig
}

func NewService(settings HashingConfig) *HashingService {
	return &HashingService{
		conf: settings,
	}
}

func (s *HashingService) generateSalt(length uint8) []byte {
	salt := make([]byte, length)
	rand.Read(salt)
	return salt
}

func (s *HashingService) HashString(str string) (string, string) {
	salt := s.generateSalt(s.conf.saltLength)
	hash := argon2.IDKey([]byte(str), salt, s.conf.time, s.conf.memory, s.conf.threads, s.conf.keyLength)
	return b64.StdEncoding.EncodeToString(hash), b64.StdEncoding.EncodeToString(salt)
}

func (s *HashingService) SameHash(str, hashedStr, salt string) (bool, error) {
	b64salt, err := b64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false, err
	}
	newPassHash := argon2.IDKey([]byte(str), b64salt, s.conf.time, s.conf.memory, s.conf.threads, s.conf.keyLength)
	newPassHashString := b64.StdEncoding.EncodeToString(newPassHash)

	return newPassHashString == hashedStr, nil
}
