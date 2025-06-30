import { faker } from '@faker-js/faker/locale/uk'
import { create } from 'zustand'
import { MessageOf } from '../components/typing-container/types/MessageInitializers'
import { MessageType } from '../components/typing-container/types/Message'

function generateRandomName() {
  return faker.hacker.adjective() + ' ' + faker.hacker.noun()
}

function generateRandomPlayerId() {
  return faker.string.uuid()
}

type MultiplayerActions = {
  setRandomName: () => void
  setRandomPlayerId: () => void
  searchForMatch: () => void
  stopSearchingForMatch: () => void
  setMatchFoundData: (data: MessageOf<MessageType.matchFound>) => void
  resetMatchFoundData: () => void
}

type MultiplayerState = {
  playerName: string
  playerId: string
  isSearchingForMatch: boolean
  matchFoundData: MessageOf<MessageType.matchFound> | null
  actions: MultiplayerActions
}

const useMultiplayerStore = create<MultiplayerState>()((set) => ({
  playerName: '',
  playerId: '',
  isSearchingForMatch: false,
  matchFoundData: null,
  actions: {
    setRandomName: () => {
      set({ playerName: generateRandomName() })
    },
    setRandomPlayerId: () => {
      set({ playerId: generateRandomPlayerId() })
    },
    searchForMatch: () => {
      set({ isSearchingForMatch: true })
    },
    stopSearchingForMatch: () => {
      set({ isSearchingForMatch: false })
    },
    setMatchFoundData: (data) => {
      set({ matchFoundData: data })
    },
    resetMatchFoundData: () => {
      set({ matchFoundData: null })
    },
  },
}))

export const usePlayerName = () =>
  useMultiplayerStore((state) => state.playerName)
export const usePlayerId = () => useMultiplayerStore((state) => state.playerId)
export const useIsSearchingForMatch = () =>
  useMultiplayerStore((state) => state.isSearchingForMatch)
export const useMatchFoundData = () =>
  useMultiplayerStore((state) => state.matchFoundData)
export const useMultiplayerActions = () =>
  useMultiplayerStore((state) => state.actions)
