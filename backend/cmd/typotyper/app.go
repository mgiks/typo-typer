package main

import (
	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/hashing"
	"github.com/mgiks/typo-typer/internal/matchmaking"
	"github.com/mgiks/typo-typer/internal/text"
	"github.com/mgiks/typo-typer/internal/token"
	"github.com/mgiks/typo-typer/internal/validation"
)

type application struct {
	config         config
	textService    text.TextService
	hashingService hashing.HashingService
	accountService account.AccountService
	tokenService   token.TokenService
	matchmaker     *matchmaking.MatchMakingService
	validator      validation.ValidationService
}

type config struct {
	port string
	db   dbConfig
}

type dbConfig struct {
	url             string
	maxConns        int32
	minIdleConns    int32
	maxConnIdleTime string
}
