package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	url  string
	pool *pgxpool.Pool
}

func New(ctx context.Context) *Database {
	dbURL := getDBURL()

	dbpool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalln("Unable to connect to database:", err)
	}

	db := Database{url: dbURL, pool: dbpool}

	if err := db.pool.Ping(ctx); err != nil {
		log.Fatalln("Unable to ping database:", err)
	}

	return &db
}

func getDBURL() string {
	requiredEnvs := map[string]string{
		"DBPass": "POSTGRES_PASSWORD",
		"DBPort": "POSTGRES_PORT",
		"DBName": "POSTGRES_DB",
	}

	envs := make(map[string]string)
	for k, v := range requiredEnvs {
		env := os.Getenv(v)

		if len(env) == 0 {
			log.Fatalf("Env `%v` not found", v)
		}

		envs[k] = env
	}

	return fmt.Sprintf(
		"postgres://postgres:%s@localhost:%s/%s",
		envs["DBPass"],
		envs["DBPort"],
		envs["DBName"],
	)
}

func (db *Database) AddText(
	ctx context.Context,
	text string,
	source string,
	uploaderName string,
) (pgx.Rows, error) {
	rows, err := db.Query(
		ctx,
		`INSERT INTO "text"(content, submitter, source) 
		VALUES ($1, $2, $3) RETURNING text, submitter`,
		text,
		uploaderName,
		source,
	)
	if err != nil {
		return nil, err
	}

	return rows, err
}

func (db *Database) GetRandomTextRow(ctx context.Context) pgx.Row {
	row := db.QueryRow(ctx, `SELECT id, content, submitter, source FROM texts ORDER BY RANDOM()`)
	return row
}

func (db *Database) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		log.Printf("Query `%v` failed: %v\n", query, err)
		return nil, err
	}

	return rows, err
}

func (db *Database) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	row := db.pool.QueryRow(ctx, query, args...)

	return row
}
