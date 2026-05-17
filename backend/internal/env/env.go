package env

import (
	"os"
	"strconv"
	"strings"
)

func GetString(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
}

func GetInt(key string, fallback int) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	valAsInt, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return fallback
	}

	return int(valAsInt)
}

func GetInt32(key string, fallback int32) int32 {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	valAsInt32, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return fallback
	}

	return int32(valAsInt32)
}

func GetStringSlice(key string, fallback []string) []string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return strings.Split(val, ",")
}

func GetByteSlice(key string, fallback []byte) []byte {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return []byte(val)
}
