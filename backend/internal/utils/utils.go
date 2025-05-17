package utils

import (
	"log"
	"os"
)

func FindEnvOrFail(env string) string {
	envVal := os.Getenv(env)
	if len(envVal) == 0 {
		log.Fatalf(`Environmental variable "%v" not found`, env)
	}
	return envVal
}
