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
	pm         *playerManager
	mm         *matchManager
	postgresDB *postgres.Database
	redisDB    *redis.Database
	wsmr       *wsMessageRouter
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
	s.pm = newPlayerManager()
	s.mm = newMatchManager()
}

func (s *server) setupDBs() {
	ctx := context.TODO()
	s.postgresDB = postgres.New(ctx)
	s.redisDB = redis.New(ctx)
}

func (s *server) setupWSRoutes() {
	s.wsmr = newWsMessageRouter()
	s.wsmr.addMessageHandler("searchForMatch", s.searchForAMatch)
}

func (s *server) setupMatchMaker() {
	go s.matchMake()
}

func (s *server) matchMake() {
	for {
		if len(s.pm.searchingPlayers) < 2 {
			continue
		}

		matchedPlayers := make([]*player, 0, 2)
		for _, player := range s.pm.searchingPlayers {
			matchedPlayers = append(matchedPlayers, player)
			delete(s.pm.searchingPlayers, player.id)
			if len(matchedPlayers) == 2 {
				break
			}
		}

		m := match{players: matchedPlayers}
		matchID := matchedPlayers[0].name + " VS " + matchedPlayers[1].name

		s.mm.mu.Lock()
		s.mm.matches[matchID] = m
		s.mm.mu.Unlock()

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
			matchedPlayers[0].name,
			matchedPlayers[1].name,
		}

		serializedMsg, err := json.Marshal(msg)
		if err != nil {
			return
		}

		for _, p := range s.mm.matches[matchID].players {
			ctx := context.TODO()
			err := p.conn.Write(ctx, websocket.MessageText, serializedMsg)

			if err != nil {
				cerr := p.conn.CloseNow()
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
