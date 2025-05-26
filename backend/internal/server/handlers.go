package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/mgiks/ttyper/internal/dtos"
	"github.com/mgiks/ttyper/internal/hashing"
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

func (s *server) postUserHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	type userData struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var ud userData

	err := json.NewDecoder(r.Body).Decode(&ud)
	if err != nil {
		log.Println("postUserHandler: failed to unmarshal request body:", err)
		http.Error(w, "", 500)
		return
	}

	salt, err := hashing.GenerateSalt()
	if err != nil {
		log.Println("postUserHandler: failed to generate salt:", err)
		http.Error(w, "", 500)
		return
	}

	ud.Password = hashing.HashAndSalt(ud.Password, salt)

	ctx := context.TODO()
	err = s.pdb.AddUserRow(ctx, ud.Name, ud.Email, ud.Password)
	if err != nil {
		log.Println("postUserHandler: failed to add user to database:", err)
		http.Error(w, "", 500)
		return
	}

	w.WriteHeader(201)
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
			log.Println(
				"websocketMessageHandler: failed to read from websocket connection:",
				err,
			)

			s.pm.mu.Lock()
			delete(s.pm.searchingPlayers, p.id)
			s.pm.mu.Unlock()

			cerr := p.conn.Close(1000, "client disconnected")
			if cerr != nil {
				log.Println(
					"websocketMessageHandler: failed to close websocket connection:",
					cerr,
				)
			}

			return
		}

		var msg dtos.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(
				"websocketMessageHandler: failed to unmarshal websocket message:",
				err,
			)
			http.Error(w, "", 500)
			return
		}

		if err = s.wsmr.routeMessage(p, msg.Type, message); err != nil {
			log.Println("websocketMessageHandler: failed to route websocket message:", err)
			http.Error(w, "", 500)
		}
	}
}
