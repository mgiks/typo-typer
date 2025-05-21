export type Message =
  | SearchForMatchMessage
  | RandomTextMessage
  | MatchFoundMessage

export enum MessageType {
  randomText = 0,
  searchForMatch,
  matchFound,
}

type RandomTextData = {
  id: number
  text: string
  submitter: string
  source: string
}

export type RandomTextMessage = {
  type: MessageType.randomText
  data: RandomTextData
}

type SearchForMatchData = {
  playerName: string
  playerId: string
}

export type SearchForMatchMessage = {
  type: MessageType.searchForMatch
  data: SearchForMatchData
}

type MatchFoundData = {
  matchID: string
  text: string
  playerNames: string[]
}

export type MatchFoundMessage = {
  type: MessageType.matchFound
  data: MatchFoundData
}
