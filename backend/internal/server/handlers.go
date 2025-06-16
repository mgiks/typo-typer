package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mgiks/ttyper/internal/auth"
	"github.com/mgiks/ttyper/internal/hashing"
	"github.com/mgiks/ttyper/internal/utils"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func (s *server) getTextHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	type textData struct {
		Id        string `json:"id"`
		Text      string `json:"text"`
		Submitter string `json:"submitter"`
		Source    string `json:"source"`
	}

	ctx := context.TODO()
	row := s.pdb.GetRandomTypingTextRow(ctx)

	var td textData
	if err := row.Scan(
		&td.Id,
		&td.Text,
		&td.Submitter,
		&td.Source,
	); err != nil {
		log.Println("getRandomTextHandler: failed to get random text row:", err)
		http.Error(w, "", 500)
		return
	}

	serializedTd, err := json.Marshal(td)
	if err != nil {
		log.Println("getRandomTextHandler: failed to marshal message:", err)
		http.Error(w, "", 500)
		return
	}

	if _, err = w.Write(serializedTd); err != nil {
		log.Println("getRandomTextHandler: failed to write message:", err)
		http.Error(w, "", 500)
	}
}

func (s *server) createPlayerHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	type userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var ud userData

	err := json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		log.Println("createPlayerHandler: failed to unmarshal request body:", err)
		http.Error(w, "", 500)
		return
	}

	salt, err := hashing.GenerateSalt()
	if err != nil {
		log.Println("createPlayerHandler: failed to generate salt:", err)
		http.Error(w, "", 500)
		return
	}

	ud.Password = hashing.Hash(ud.Password, salt)

	ctx := context.TODO()
	err = s.pdb.AddPlayerRow(ctx, ud.Name, ud.Email, ud.Password)
	if err != nil {
		log.Println("createPlayerHandler: failed to add user to database:", err)
		http.Error(w, "", 500)
		return
	}

	w.WriteHeader(201)
}

func (s *server) signInHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	type userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var ud userData

	err := json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		log.Println("signInHandler: failed to unmarshal request body:", err)
		http.Error(w, "", 500)
		return
	}

	ctx := context.TODO()
	pr := s.pdb.GetPlayerRow(ctx, ud.Name, ud.Email)

	var hashedPass string
	err = pr.Scan(&hashedPass)
	if err != nil {
		log.Println("signInHandler: failed to scan player row:", err)
		http.Error(w, "", 500)
		return
	}

	isCorrect, err := hashing.IsEqualToHash(ud.Password, hashedPass)
	if err != nil {
		log.Println("signInHandler: failed to compare password to hashed password:", err)
		http.Error(w, "", 500)
		return
	}

	if !isCorrect {
		http.Error(w, "Authenticatio failed: Invalid username or password", http.StatusUnauthorized)
		return
	}

	seed := []byte(utils.FindEnvOr("privatekeyseed", func() {
		log.Fatal("signInHandler: failed to find 'privatekeyseed' env")
	}))

	expiresAt := time.Now().UTC().Add(time.Minute * 15)
	jwt, err := auth.GenerateJWT(seed, jwt.MapClaims{
		"player": ud.Name,
		"nbf":    expiresAt.Unix(),
	})
	if err != nil {
		log.Println("signInHandler: failed to generate jwt:", err)
		http.Error(w, "", 500)
		return
	}

	type jwtData struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	data, err := json.Marshal(jwtData{jwt, "bearer", int64(expiresAt.Sub(time.Now().UTC()).Seconds())})
	if err != nil {
		log.Println("signInHandler: failed to marshal json:", err)
		http.Error(w, "", 500)
		return
	}

	if _, err := w.Write(data); err != nil {
		log.Println("signInHandler: failed to write message:", err)
		http.Error(w, "", 500)
	}
}

var acceptOptions = &websocket.AcceptOptions{
	OriginPatterns: []string{"localhost:5173"},
}

func (s *server) websocketMessageHandler(
	w http.ResponseWriter,
	r *http.Request,
) {
	wsc, err := websocket.Accept(w, r, acceptOptions)
	if err != nil {
		log.Println("websocketMessageHandler: failed to establish websocket connection:", err)
		http.Error(w, "failed to upgrade to websocket connection", 400)
		return
	}

	p := &player{conn: wsc}
	ctx := context.TODO()
	for {
		messageType, message, err := p.conn.Read(ctx)
		if err != nil || messageType != websocket.MessageText {
			log.Println("websocketMessageHandler: failed to read from websocket connection:", err)

			s.pm.mu.Lock()
			delete(s.pm.searchingPlayers, p.id)
			s.pm.mu.Unlock()

			cerr := p.conn.Close(1000, "client disconnected")
			if cerr != nil {
				log.Println("websocketMessageHandler: failed to close websocket connection:", cerr)
			}

			return
		}

		var msg wsMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("websocketMessageHandler: failed to unmarshal websocket message:", err)
			http.Error(w, "", 500)
			return
		}

		if err = s.wsmr.routeMessage(p, msg.Type, message); err != nil {
			log.Println("websocketMessageHandler: failed to route websocket message:", err)
			http.Error(w, "", 500)
		}
	}
}
