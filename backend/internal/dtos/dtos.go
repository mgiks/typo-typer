package dtos

type randomTextData struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	Submitter string `json:"submitter"`
	Source    string `json:"source"`
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

type Message struct {
	Type string `json:"type"`
}

type RandomTextMessage struct {
	Message
	Data randomTextData `json:"data"`
}

type MatchFoundMessage struct {
	Message
	Data matchFoundData `json:"data"`
}

type SearchForMatchMessage struct {
	Message
	Data searchForMatchData `json:"data"`
}

func NewMatchFoundMessage() *MatchFoundMessage {
	msg := &MatchFoundMessage{}
	msg.Type = "matchFound"
	return msg
}

func NewRandomTextMessage() *RandomTextMessage {
	msg := &RandomTextMessage{}
	msg.Type = "randomText"
	return msg
}
