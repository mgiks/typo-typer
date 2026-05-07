package main

import (
	"context"
	"log"

	"github.com/mgiks/typo-typer/internal/db"
	"github.com/mgiks/typo-typer/internal/env"
	"github.com/mgiks/typo-typer/internal/storage"
)

func main() {
	dbUrl := env.GetString("DB_URL", "postgres://admin:adminpassword@localhost:5433/typo-typer")
	conn, err := db.New(context.Background(), dbUrl, 1, 1, "15s")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	store := storage.NewStore(conn)
	if err := db.Seed(store); err != nil {
		log.Fatal("database seeding failed", err.Error())
	}
}
