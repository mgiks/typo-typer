package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/database"
	"github.com/mgiks/typo-typer/internal/handler"
	"github.com/mgiks/typo-typer/internal/hashing"
)

const port = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", "error", err)
		return
	}

	ctx := context.Background()
	db, err := database.Connect(ctx)
	if err != nil {
		slog.Error("database connection failed", "error", err)
		return
	}

	hashingService := hashing.NewService(hashing.DefaultHashingConfig)
	accountService := account.NewService(db, hashingService)
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/texts", handler.NewGetTextHandler(db))
	mux.HandleFunc("POST /auth/register", handler.NewRegisterHandler(accountService, validate))
	mux.HandleFunc("POST /auth/login", handler.NewLoginHandler(accountService, validate))

	fmt.Printf("Listening and serving on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		slog.Error("http listening and serving failed", "error", err)
	}

}
