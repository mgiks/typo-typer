package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/mgiks/ttyper/internal/db/postgres"
	"github.com/mgiks/ttyper/internal/db/redis"
)

type server struct {
	mux  http.ServeMux
	pm   *playerManager
	mm   *matchManager
	pdb  *postgres.Database
	rdb  *redis.Database
	wsmr *wsMessageRouter
}

func New() *server {
	s := server{}

	s.setupManagers()
	s.setupDBs()
	s.setupWSRoutes()
	s.setupMatchMaker()

	s.mux.HandleFunc("/", s.websocketMessageHandler)
	s.mux.HandleFunc("GET /texts", s.getTextHandler)
	s.mux.HandleFunc("POST /players", s.createPlayerHandler)
	s.mux.HandleFunc("POST /auth/signin", s.signInHandler)

	return &s
}

func (s *server) setupManagers() {
	s.pm = newPlayerManager()
	s.mm = newMatchManager()
}

func (s *server) setupDBs() {
	ctx := context.TODO()
	s.pdb = postgres.New(ctx)
	s.rdb = redis.New(ctx)
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
		row := s.pdb.GetRandomTypingTextRow(ctx)

		var text string
		if err := row.Scan(nil, &text); err != nil {
			log.Println("matchMake: failed to get random text row:", err)
			continue
		}

		msg := initializeMatchFoundMessage()
		msg.Data.MatchID = matchID
		msg.Data.Text = text
		msg.Data.PlayerNames = []string{
			matchedPlayers[0].name,
			matchedPlayers[1].name,
		}

		serializedMsg, err := json.Marshal(msg)
		if err != nil {
			log.Println("matchMake: failed to marshal random text row:", err)
			continue
		}

		for _, p := range s.mm.matches[matchID].players {
			ctx := context.TODO()
			err := p.conn.Write(ctx, websocket.MessageText, serializedMsg)
			if err != nil {
				log.Println("matchMake: failed to write message")

				cerr := p.conn.CloseNow()
				if cerr != nil {
					log.Println("matchMake: failed to close websocket connection:", cerr)
				}
			}
		}
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
