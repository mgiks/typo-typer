package handler

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mgiks/typo-typer/internal/matchmaking"
)

type matchEnterMsg struct {
	Name string `json:"name"`
}

type matchMaker interface {
	EnterMatch(matchId string, p matchmaking.MatchedPlayer) error
	MatchExists(matchId string) bool
	PlayerBelongs(matchId, playerId string) bool
}

type accessTokenParser interface {
	ParseAccessToken(ctx context.Context, tokenString string) (jwt.MapClaims, error)
}

// TODO: Handle authorization and figure out the layer separation when doing that
func NewEnterMatchHandler(mm matchMaker, tp accessTokenParser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchId := r.PathValue("matchId")
		if len(matchId) == 0 {
			writeJSON(w, http.StatusBadRequest, apiErr{Error: "match id not provided"})
			return
		}

		t, err := r.Cookie("access_token")
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, apiErr{Error: "no 'access_token' cookie provided"})
			return
		}

		claims, err := tp.ParseAccessToken(r.Context(), t.Value)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, apiErr{Error: "failed to parse access token"})
			return
		}

		playerId, err := claims.GetSubject()
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, apiErr{Error: "failed to get 'sub' from access token"})
			return
		}

		if !mm.MatchExists(matchId) {
			writeJSON(w, http.StatusUnauthorized, apiErr{Error: "match does not exist"})
			return
		}

		if !mm.PlayerBelongs(matchId, playerId) {
			writeJSON(w, http.StatusUnauthorized, apiErr{Error: "not authorized to enter the match"})
			return
		}

		conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
			OriginPatterns: []string{os.Getenv("FRONTEND_URL")},
		})
		if err != nil {
			slog.Error("match entering websocket handshake failed", "error", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var msg matchEnterMsg
		err = wsjson.Read(ctx, conn, &msg)
		if err != nil {
			slog.Error("match entering websocket message reading failed", "error", err)
			return
		}

		p := matchmaking.NewMatchedPlayer(msg.Name, conn)
		if err = mm.EnterMatch(matchId, p); err != nil {
			slog.Error("failed to enter match:", "error", err)
			conn.Close(websocket.StatusInternalError, "internal server error")
		}
	}
}
