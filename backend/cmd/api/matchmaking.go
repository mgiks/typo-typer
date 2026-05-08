package main

import (
	"errors"
	"net/http"

	"github.com/coder/websocket"
	"github.com/go-chi/chi/v5"
	"github.com/mgiks/typo-typer/internal/matchmaker"
)

type JoinPoolMsg struct {
	Username string `json:"username" validate:"required,max=30"`
	WPM      uint16 `json:"password" validate:"required"`
}

func (app application) joinPoolHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := app.acceptWSHandshake(w, r, r.URL.Path)
	if err != nil {
		return
	}

	defer conn.Close(websocket.StatusNormalClosure, "")

	var msg JoinPoolMsg
	if err := readJSON(w, r, &msg); err != nil {
		conn.Close(StatusInavalidMessage, "message decoding failed")
		return
	}

	if err := app.validator.ValidateJSON(msg); err != nil {
		conn.Close(StatusInavalidMessage, "invalid message")
		return
	}

	p := matchmaker.NewSearchingPlayer(msg.Username, msg.WPM, conn)
	app.matchmaker.JoinPool(p)
}

type EnterMatchMsg struct {
	Username string `json:"name" validate:"max=30"`
}

func (app application) enterMatchHandler(w http.ResponseWriter, r *http.Request) {
	matchID := chi.URLParam(r, "matchID")

	t, err := r.Cookie("access_token")
	if err != nil {
		app.unauthorized(w, r, errors.New("no 'access_token' cookie provided"))
		return
	}

	claims, err := app.tokenService.ParseAccessToken(r.Context(), t.Value)
	if err != nil {
		app.unauthorized(w, r, errors.New("failed ot parse access token"))
		return
	}

	playerId, err := claims.GetSubject()
	if err != nil {
		app.unauthorized(w, r, errors.New("failed to get 'sub' from access token"))
		return
	}

	if !app.matchmaker.MatchExists(matchID) {
		app.badRequestResponse(w, r, errors.New("match does not exist"))
		return
	}

	if !app.matchmaker.PlayerBelongs(matchID, playerId) {
		writeJSON(w, http.StatusUnauthorized, errors.New("not authorized to enter the match"))
		return
	}

	conn, err := app.acceptWSHandshake(w, r, r.URL.Path)
	if err != nil {
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	var msg EnterMatchMsg
	if err := readJSON(w, r, &msg); err != nil {
		conn.Close(StatusInavalidMessage, "message decoding failed")
		return
	}

	if err := app.validator.ValidateJSON(msg); err != nil {
		conn.Close(StatusInavalidMessage, "invalid message")
		return
	}

	p := matchmaker.NewMatchedPlayer(msg.Username, conn)
	if err = app.matchmaker.EnterMatch(matchID, p); err != nil {
		conn.Close(websocket.StatusInternalError, "internal server error")
	}
}
