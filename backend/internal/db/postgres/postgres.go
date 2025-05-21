package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mgiks/ttyper/internal/utils"
)

type Database struct {
	*pgxpool.Pool
}

func New(ctx context.Context) *Database {
	dbURL := getDBURL()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalln("unable to connect to postgres database:", err)
	}

	db := Database{dbpool}

	if err := db.Ping(ctx); err != nil {
		log.Fatalln("unable to ping postgres database:", err)
	}

	return &db
}

func getDBURL() string {
	dbPass := utils.FindEnvOrFail("POSTGRES_PASSWORD")
	dbPort := utils.FindEnvOrFail("POSTGRES_PORT")
	dbName := utils.FindEnvOrFail("POSTGRES_DB")

	return fmt.Sprintf(
		"postgres://postgres:%s@localhost:%s/%s",
		dbPass,
		dbPort,
		dbName,
	)
}

func (db *Database) AddTypingTextRow(
	ctx context.Context,
	text string,
	source string,
	uploaderName string,
) error {
	_, err := db.Query(
		ctx,
		`INSERT INTO typing_text(content, submitter, source) 
		VALUES ($1, $2, $3) RETURNING text, submitter`,
		text,
		uploaderName,
		source,
	)
	return err
}

func (db *Database) GetRandomTypingTextRow(ctx context.Context) pgx.Row {
	return db.QueryRow(
		ctx,
		`SELECT id, content, submitter, source 
		FROM typing_text ORDER BY RANDOM()`,
	)
}
