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
	const (
		pass = "POSTGRES_PASSWORD"
		port = "POSTGRES_PORT"
		db   = "POSTGRES_DB"
	)

	makeCallback := func(key string) func() {
		return func() {
			log.Fatalf("getDBURL: failed to find %q environmental variable", key)
		}
	}

	dbPass := utils.FindEnvOr(pass, makeCallback(pass))
	dbPort := utils.FindEnvOr(port, makeCallback(port))
	dbName := utils.FindEnvOr(db, makeCallback(db))

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
		VALUES ($1, $2, $3)`,
		text,
		uploaderName,
		source,
	)
	if err != nil {
		return fmt.Errorf("AddTypingTextRow: failed to add row: %w", err)
	}

	return nil
}

func (db *Database) AddUserRow(
	ctx context.Context,
	name string,
	email string,
	hashedPassword string,
) error {
	_, err := db.Query(
		ctx,
		`INSERT INTO "user"(name, email, hashed_password)
		VALUES ($1, $2, $3)`,
		name, email, hashedPassword,
	)
	if err != nil {
		return fmt.Errorf("AddUserRow: failed to add row: %w", err)
	}

	return nil
}

func (db *Database) GetRandomTypingTextRow(ctx context.Context) pgx.Row {
	return db.QueryRow(
		ctx,
		`SELECT id, content, submitter, source 
		FROM typing_text ORDER BY RANDOM()`,
	)
}
