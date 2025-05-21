export type MessageDataMap = {
  [MessageType.randomText]: {
    id: number
    text: string
    submitter: string
    source: string
  }
  [MessageType.searchForMatch]: {
    playerName: string
    playerId: string
  }
  [MessageType.matchFound]: {
    matchID: string
    text: string
    playerNames: string[]
  }
}

export enum MessageType {
  randomText = 'randomText',
  searchForMatch = 'searchForMatch',
  matchFound = 'matchFound',
}

export type Message = {
  [T in keyof MessageDataMap]: {
    type: T
    data: MessageDataMap[T]
  }
}[keyof MessageDataMap]
