package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/mgiks/ttyper/internal/dtos"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
}

func (s *server) getRandomTextHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	ctx := context.TODO()
	row := s.postgresDB.GetRandomTypingTextRow(ctx)

	msg := dtos.InitializeRandomTextMessage()
	if err := row.Scan(
		&msg.Data.Id,
		&msg.Data.Text,
		&msg.Data.Submitter,
		&msg.Data.Source,
	); err != nil {
		log.Println("getRandomTextHandler: failed to get random text row:", err)
		http.Error(w, "", 500)
		return
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		log.Println("getRandomTextHandler: failed to marshal message:", err)
		http.Error(w, "", 500)
		return
	}

	if _, err = w.Write(serializedMsg); err != nil {
		log.Println("getRandomTextHandler: failed to write message:", err)
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
