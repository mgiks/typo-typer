package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/mgiks/ttyper/internal/db/postgres"
	"github.com/mgiks/ttyper/internal/db/redis"
	"github.com/mgiks/ttyper/internal/dtos"
)

type server struct {
	mux        http.ServeMux
	pm         *PlayerManager
	mm         *MatchManager
	postgresDB *postgres.Struct
	redisDB    *redis.Struct
	wsmr       *WSMessageRouter
}

func New() *server {
	s := server{}

	s.setupManagers()
	s.setupDBs()
	s.setupWSRoutes()
	s.setupMatchMaker()

	s.mux.HandleFunc("/", s.websocketMessageHandler)
	s.mux.HandleFunc("GET /random-texts", s.getRandomTextHandler)

	return &s
}

func (s *server) setupManagers() {
	s.pm = NewPlayerManager()
	s.mm = NewMatchManager()
}

func (s *server) setupDBs() {
	ctx := context.TODO()
	s.postgresDB = postgres.New(ctx)
	s.redisDB = redis.New(ctx)
}

func (s *server) setupWSRoutes() {
	s.wsmr = NewWSMessageRouter()
	s.wsmr.AddMessageHandler("searchForMatch", s.SearchForAMatch)
}

func (s *server) setupMatchMaker() {
	go s.matchMake()
}

func (s *server) matchMake() {
	for {
		if len(s.pm.SearchingPlayers) < 2 {
			continue
		}

		matchedPlayers := make([]*Player, 0, 2)
		for _, player := range s.pm.SearchingPlayers {
			matchedPlayers = append(matchedPlayers, player)
			delete(s.pm.SearchingPlayers, player.Id)
			if len(matchedPlayers) == 2 {
				break
			}
		}

		m := Match{Players: matchedPlayers}
		matchID := matchedPlayers[0].Name + " VS " + matchedPlayers[1].Name

		s.mm.Mu.Lock()
		s.mm.Matches[matchID] = m
		s.mm.Mu.Unlock()

		ctx := context.TODO()
		row := s.postgresDB.GetRandomTextRow(ctx)

		var text string
		if err := row.Scan(nil, &text); err != nil {
			log.Println("Failed to get random text row:", err)
		}

		msg := dtos.NewMatchFoundMessage()
		msg.Data.MatchID = matchID
		msg.Data.Text = text
		msg.Data.PlayerNames = []string{
			matchedPlayers[0].Name,
			matchedPlayers[1].Name,
		}

		serializedMsg, err := json.Marshal(msg)
		if err != nil {
			return
		}

		for _, p := range s.mm.Matches[matchID].Players {
			ctx := context.TODO()
			err := p.Conn.Write(ctx, websocket.MessageText, serializedMsg)

			if err != nil {
				cerr := p.Conn.CloseNow()
				if cerr != nil {
					log.Println("Failed to close websocket connection:", cerr)
				}
			}
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
