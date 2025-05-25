package dtos

type matchFoundData struct {
	MatchID     string   `json:"matchID"`
	Text        string   `json:"text"`
	PlayerNames []string `json:"playerNames"`
}

type searchForMatchData struct {
	PlayerName string `json:"playerName"`
	PlayerId   string `json:"playerId"`
}

type Message struct {
	Type string `json:"type"`
}

type MatchFoundMessage struct {
	Message
	Data matchFoundData `json:"data"`
}

type SearchForMatchMessage struct {
	Message
	Data searchForMatchData `json:"data"`
}

func InitializeMatchFoundMessage() *MatchFoundMessage {
	msg := &MatchFoundMessage{}
	msg.Type = "matchFound"
	return msg
}
