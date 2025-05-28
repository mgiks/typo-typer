package server

type wsMessage struct {
	Type string `json:"type"`
}

type matchFoundData struct {
	MatchID     string   `json:"matchID"`
	Text        string   `json:"text"`
	PlayerNames []string `json:"playerNames"`
}

type searchForMatchData struct {
	PlayerName string `json:"playerName"`
	PlayerId   string `json:"playerId"`
}

type matchFoundMessage struct {
	wsMessage
	Data matchFoundData `json:"data"`
}

type searchForMatchMessage struct {
	wsMessage
	Data searchForMatchData `json:"data"`
}

func initializeMatchFoundMessage() *matchFoundMessage {
	return &matchFoundMessage{wsMessage: wsMessage{Type: "matchFound"}}
}
