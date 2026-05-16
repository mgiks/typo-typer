package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mgiks/typo-typer/internal/storage"
	"github.com/mgiks/typo-typer/internal/user"
)

type RegisterPayload struct {
	Username string `json:"username" validate:"required,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}

func (app application) registerHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to read json: %w", err))
		return
	}

	if err := app.validator.ValidateJSON(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	acc := storage.User{
		Username: payload.Username,
		Email:    payload.Email,
	}
	acc.Password.SetText(payload.Username)

	if err := app.userService.CreateUser(r.Context(), &acc); err != nil {
		switch {
		case errors.Is(err, user.ErrUserAlreadyExists),
			errors.Is(err, user.ErrUsernameEmpty),
			errors.Is(err, user.ErrPasswordTooShort),
			errors.Is(err, user.ErrIncorrectPassword):
			app.badRequestResponse(w, r, err)
		case errors.Is(err, user.ErrUserNotFound):
			app.notFoundResponse(w, r, err)
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
		app.badRequestResponse(w, r, err)
		return
	}

	if err := app.userService.PasswordCorrect(r.Context(), payload.Username, payload.Password); err != nil {
		switch {
		case errors.Is(err, user.ErrUserNotFound),
			errors.Is(err, user.ErrIncorrectPassword):
			app.badRequestResponse(w, r, err)
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
