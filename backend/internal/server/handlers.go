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
	row := s.postgresDB.GetRandomTextRow(ctx)

	msg := dtos.NewRandomTextMessage()
	if err := row.Scan(
		&msg.Data.Id,
		&msg.Data.Text,
		&msg.Data.Submitter,
		&msg.Data.Source,
	); err != nil {
		log.Println("Failed to get random text row:", err)
		return
	}

	serializedMsg, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(serializedMsg); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
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
		cerr := wsc.CloseNow()
		if cerr != nil {
			log.Println("Failed to close websocket connection:", cerr)
		}
		return
	}

	p := &Player{Conn: wsc}
	ctx := context.TODO()
	for {
		messageType, message, err := p.Conn.Read(ctx)
		if err != nil || messageType != websocket.MessageText {
			s.pm.Mu.Lock()
			delete(s.pm.SearchingPlayers, p.Id)
			s.pm.Mu.Unlock()

			cerr := p.Conn.CloseNow()
			if cerr != nil {
				log.Println("Failed to close websocket connection:", cerr)
				return
			}
		}

		var msg dtos.Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Failed to unmarshal websocketMessage:", err)
		}

		s.wsmr.RouteMessage(p, msg.Type, message)
	}
}
