package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/coder/websocket"
	"github.com/mgiks/ttyper/internal/db"
	"github.com/mgiks/ttyper/internal/dtos"
)

type server struct {
	mux  http.ServeMux
	pm   *PlayerManager
	mm   *MatchManager
	db   *db.Database
	wsmr *WSMessageRouter
}

func New() *server {
	s := server{}

	s.setupManagers()
	s.setupDB()
	s.setupWSRoutes()
	s.setupMatchmaker()

	s.mux.HandleFunc("/", s.websocketMessageHandler)
	s.mux.HandleFunc("GET /random-texts", s.getRandomTextHandler)

	return &s
}

func (s *server) setupManagers() {
	s.pm = NewPlayerManager()
	s.mm = NewMatchManager()
}

func (s *server) setupDB() {
	ctx := context.TODO()
	s.db = db.New(ctx)
}

func (s *server) setupWSRoutes() {
	s.wsmr = NewWSMessageRouter()
	s.wsmr.AddMessageHandler("searchForMatch", s.SearchForAMatch)
}

func (s *server) setupMatchmaker() {
	go s.matchPlayers()
}

func (s *server) matchPlayers() {
	for {
		if len(s.pm.SearchingPlayers) >= 2 {
			matchedPlayers := make([]*Player, 0, 2)
			for _, v := range s.pm.SearchingPlayers {
				matchedPlayers = append(matchedPlayers, v)
				if len(matchedPlayers) == 2 {
					break
				}
			}
			playerIDs := make([]string, 0, 2)
			playerNames := make([]string, 0, 2)
			for _, p := range matchedPlayers {
				fmt.Println("Deleting", p.Name)
				playerIDs = append(playerIDs, p.Id)
				playerNames = append(playerNames, p.Name)
				delete(s.pm.SearchingPlayers, p.Id)
			}
			m := Match{Players: matchedPlayers}
			matchID := strings.Join(playerIDs, " VS ")
			s.mm.Mu.Lock()
			s.mm.Matches[matchID] = m
			s.mm.Mu.Unlock()

			ctx := context.TODO()
			row := s.db.GetRandomTextRow(ctx)

			var text string
			if err := row.Scan("", &text); err != nil {
				log.Println("Failed to get random text row:", err)
			}

			mfm := dtos.NewMatchFoundMessage()
			mfm.Data.MatchID = matchID
			mfm.Data.Text = text
			mfm.Data.PlayerNames = playerNames

			jsonMfm, err := json.Marshal(mfm)
			if err != nil {
				return
			}

			for _, p := range s.mm.Matches[matchID].Players {
				err := p.Conn.Write(context.TODO(), websocket.MessageText, jsonMfm)
				fmt.Println(mfm)
				if err != nil {
					cerr := p.Conn.CloseNow()
					if cerr != nil {
						log.Println("Failed to close websocket connection:", cerr)
					}
				}
			}
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
