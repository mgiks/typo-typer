package message

type matchFoundData struct {
	MatchID     string   `json:"matchID"`
	Text        string   `json:"text"`
	PlayerNames []string `json:"playerNames"`
}

type searchForMatchData struct {
	PlayerName string `json:"playerName"`
	PlayerId   string `json:"playerId"`
}

type Struct struct {
	Type string `json:"type"`
}

type MatchFound struct {
	Struct
	Data matchFoundData `json:"data"`
}

type SearchForMatch struct {
	Struct
	Data searchForMatchData `json:"data"`
}

func InitializeMatchFound() *MatchFound {
	return &MatchFound{Struct: Struct{Type: "matchFound"}}
}
