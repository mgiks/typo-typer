package utils

import (
	"os"
)

func FindEnvOr(key string, callback func()) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		callback()
	}

	return val
}
