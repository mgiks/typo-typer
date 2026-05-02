package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
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
