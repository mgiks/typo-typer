export type MessageDataMap = {
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
  searchForMatch = 'searchForMatch',
  matchFound = 'matchFound',
}

export type Message = {
  [T in keyof MessageDataMap]: {
    type: T
    data: MessageDataMap[T]
  }
}[keyof MessageDataMap]
