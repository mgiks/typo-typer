package utils

import (
	"log"
	"os"
)

func FindEnvOrFail(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf(`Environmental variable "%v" not found`, key)
	}

	return val
}
