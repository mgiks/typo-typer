import { createSlice } from '@reduxjs/toolkit'

export type PlayerStatusState = {
  isTyping: boolean
  finishedTyping: boolean
}

export const playerStatusInitialState: PlayerStatusState = {
  isTyping: false,
  finishedTyping: false,
}

export const playerStatusSlice = createSlice({
  name: 'playerStatus',
  initialState: playerStatusInitialState,
  reducers: {
    playerIsTyping: (state) => {
      state.isTyping = true
    },
    playerIsNotTyping: (state) => {
      state.isTyping = false
    },
    playerFinishedTyping: (state) => {
      state.finishedTyping = true
      state.isTyping = false
    },
    resetPlayerStatus: (state) => {
      state.isTyping = playerStatusInitialState.isTyping
      state.finishedTyping = playerStatusInitialState.finishedTyping
    },
  },
})

export const {
  playerIsTyping,
  playerIsNotTyping,
  playerFinishedTyping,
  resetPlayerStatus,
} = playerStatusSlice.actions

export default playerStatusSlice.reducer
