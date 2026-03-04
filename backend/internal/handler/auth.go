package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mgiks/typo-typer/internal/account"
)

type registrationData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type registrationResponse struct {
	Username string `json:"username"`
}

type AccountCreator interface {
	CreateAccount(ctx context.Context, username, password string) error
}

type Validator interface {
	Struct(s interface{}) error
}

func NewRegisterHandler(c AccountCreator, v Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data registrationData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			slog.Error("register request data decoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		err := v.Struct(data)
		if validationErr(err, w, "registration data validation failed") {
			return
		}

		err = c.CreateAccount(r.Context(), data.Username, data.Password)
		if serviceErr(err, w, "account creation failed") {
			return
		}

		w.WriteHeader(http.StatusCreated)
		resp := registrationResponse{Username: data.Username}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			slog.Error("registration response encoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}
	}
}

type loginData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	Username string `json:"username"`
}

type AccountPasswordChecker interface {
	PasswordCorrect(ctx context.Context, username, password string) error
}

func NewLoginHandler(c AccountPasswordChecker, v Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var data loginData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			slog.Error("login request decoding failed", "error", err)
			writeInternalServerErrorJSON(w)
			return
		}

		err := v.Struct(data)
		if validationErr(err, w, "login data validation failed") {
			return
		}

		err = c.PasswordCorrect(r.Context(), data.Username, data.Password)
		if serviceErr(err, w, "login password check failed") {
			return
		}

		resp := loginResponse{Username: data.Username}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			slog.Error("login response encoding failed", "error", err)
			writeInternalServerErrorJSON(w)
		}
	}
}

func validationErr(err error, w http.ResponseWriter, logMsg string) bool {
	if err == nil {
		return false
	}

	var validateErrs validator.ValidationErrors
	if errors.As(err, &validateErrs) {
		writeJSON(w, http.StatusBadRequest, APIErrDetailed{
			Error:   "incorrect json data",
			Details: buildValidationErrJSON(validateErrs)},
		)
	} else {
		slog.Error(logMsg, "error", err)
		writeInternalServerErrorJSON(w)
	}

	return true
}

func serviceErr(err error, w http.ResponseWriter, logMsg string) bool {
	if err == nil {
		return false
	}

	switch {
	case errors.Is(err, account.ErrAccountAlreadyExists),
		errors.Is(err, account.ErrUsernameEmpty),
		errors.Is(err, account.ErrPasswordTooShort):
		writeJSON(w, http.StatusBadRequest, APIErr{Error: err.Error()})
	case errors.Is(err, account.ErrAccountNotFound),
		errors.Is(err, account.ErrIncorrectPassword):
		writeJSON(w, http.StatusBadRequest, APIErr{Error: "incorrect username or password"})
	default:
		slog.Error(logMsg, "error", err)
		writeInternalServerErrorJSON(w)
	}

	return true
}

type fieldErr struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

func buildValidationErrJSON(validateErrs validator.ValidationErrors) []fieldErr {
	var errJson []fieldErr

	for _, v := range validateErrs {
		errJson = append(errJson, fieldErr{
			Field: v.Field(),
			Rule:  v.Tag(),
		})
	}

	return errJson
}
