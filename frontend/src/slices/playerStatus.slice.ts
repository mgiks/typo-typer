import { createSlice } from '@reduxjs/toolkit'

export type playerStatusState = {
  startedTyping: boolean
  finishedTyping: boolean
}

export const playerStatusInitialState: playerStatusState = {
  startedTyping: false,
  finishedTyping: false,
}

export const playerStatusSlice = createSlice({
  name: 'playerStatus',
  initialState: playerStatusInitialState,
  reducers: {
    playerStartedTyping: (state) => {
      state.startedTyping = true
    },
    playerFinishedTyping: (state) => {
      state.finishedTyping = true
    },
  },
})

export const { playerStartedTyping, playerFinishedTyping } =
  playerStatusSlice.actions

export default playerStatusSlice.reducer
