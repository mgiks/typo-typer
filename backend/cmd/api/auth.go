package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mgiks/typo-typer/internal/account"
	"github.com/mgiks/typo-typer/internal/storage"
)

type RegisterPayload struct {
	Username string `json:"username" validate:"required,max=30"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
}

func (app application) registerHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to read json: %w", err))
		return
	}

	if err := app.validator.ValidateJSON(payload); err != nil {
		app.badRequest(w, r, fmt.Errorf("invalid json: %w", err))
		return
	}

	acc := storage.Account{
		Username: payload.Username,
		Password: payload.Password,
		Email:    payload.Email,
	}

	if err := app.accountService.CreateAccount(r.Context(), &acc); err != nil {
		switch {
		case errors.Is(err, account.ErrAccountAlreadyExists),
			errors.Is(err, account.ErrUsernameEmpty),
			errors.Is(err, account.ErrPasswordTooShort),
			errors.Is(err, account.ErrIncorrectPassword),
			errors.Is(err, account.ErrAccountNotFound):
			app.badRequest(w, r, err)
		default:
			app.internalServerError(w, r, fmt.Errorf("failed to create account: %w", err))
		}
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, acc); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to send json response: %w", err))
	}
}

type LoginPayload struct {
	Username string `json:"username" validate:"required,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

func (app application) loginHandler(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to read json: %w", err))
		return
	}

	if err := app.validator.ValidateJSON(payload); err != nil {
		app.badRequest(w, r, fmt.Errorf("invalid json: %w", err))
		return
	}

	if err := app.accountService.PasswordCorrect(r.Context(), payload.Username, payload.Password); err != nil {
		switch {
		case errors.Is(err, account.ErrAccountNotFound),
			errors.Is(err, account.ErrIncorrectPassword):
			app.badRequest(w, r, err)
		default:
			app.internalServerError(w, r, fmt.Errorf("failed to verify password: %w", err))
		}
		return
	}

	accessToken, err := app.tokenService.CreateAccessToken(r.Context(), payload.Username)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to create access token"))
		return
	}

	refreshToken, err := app.tokenService.CreateRefreshToken(r.Context(), payload.Username)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to create refresh token"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/matchmaking",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusNoContent)
}
