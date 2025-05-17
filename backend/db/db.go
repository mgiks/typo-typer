package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mgiks/ttyper/hashing"
)

type Database struct {
	Url     string
	Pool    *pgxpool.Pool
	Context context.Context
}

func ConnectToDB(ctx context.Context) *Database {
	envs := map[string]string{
		"dbPass": "POSTGRES_PASSWORD",
		"dbPort": "POSTGRES_PORT",
		"dbName": "POSTGRES_DB",
	}

	foundEnvs := make(map[string]string)

	for k, v := range envs {
		env := os.Getenv(v)

		if len(env) == 0 {
			log.Fatalf("Env `%v` not found", v)
		}

		foundEnvs[k] = env
	}

	dbPass := foundEnvs["dbPass"]
	dbPort := foundEnvs["dbPort"]
	dbName := foundEnvs["dbName"]

	dbUrl := fmt.Sprintf(
		"postgres://postgres:%s@localhost:%s/%s",
		dbPass,
		dbPort,
		dbName,
	)

	dbPool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to Database: %v\n", err)
	}

	db := Database{Url: dbUrl, Pool: dbPool, Context: ctx}

	err = db.Pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping Database: %v\n", err)
	}

	return &db
}

func (db *Database) Query(query string, args ...any) (pgx.Rows, error) {
	rows, err := db.Pool.Query(db.Context, query, args...)
	if err != nil {
		log.Printf("Query `%v` failed: %v\n", query, err)
		return nil, err
	}

	return rows, err
}

func (db *Database) QueryRow(query string, args ...any) pgx.Row {
	row := db.Pool.QueryRow(db.Context, query, args...)

	return row
}

func (db *Database) AddText(text string, uploaderName string) (pgx.Rows, error) {
	rows, err := db.Query(
		`INSERT INTO "text"(content, uploader_name) 
		VALUES ($1, $2) RETURNING text, uploader_name`,
		text,
		uploaderName,
	)
	if err != nil {
		return nil, err
	}

	return rows, err
}

func (db *Database) AddUser(name string, email string, password string) (pgx.Rows, error) {
	salt, err := hashing.GenerateSalt()
	if err != nil {
		return nil, err
	}

	hashedPassword := hashing.HashAndSalt(password, salt)
	rows, err := db.Query(
		`INSERT INTO "user"(username, email, password)
		VALUES ($1, $2, $3) RETURNING username, email`,
		name,
		email,
		hashedPassword,
	)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (db *Database) RandomTextRow() pgx.Row {
	row := db.QueryRow(`SELECT id, content, submitter, source FROM "text" ORDER BY RANDOM()`)
	return row
}

func (db *Database) RandomText() string {
	row := db.QueryRow(`SELECT content FROM "text" ORDER BY RANDOM()`)
	var text string
	err := row.Scan(&text)
	if err != nil {
		log.Printf("Failed to get random text: %v", err)
	}
	return text
}
