package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/mgiks/ttyper/internal/utils"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	client *redis.Client
}

func New(ctx context.Context) *Database {
	opt, err := redis.ParseURL(getDBURL())
	if err != nil {
		log.Fatalln("Unable to connect to redis database:", err)
	}

	return &Database{client: redis.NewClient(opt)}
}

func getDBURL() string {
	dbPass := utils.FindEnvOrFail("REDIS_PASSWORD")
	dbPort := utils.FindEnvOrFail("REDIS_PORT")
	dbName := utils.FindEnvOrFail("REDIS_DB")

	return fmt.Sprintf(
		"redis://redis:%s@localhost:%s/%s",
		dbPass,
		dbPort,
		dbName,
	)
}
