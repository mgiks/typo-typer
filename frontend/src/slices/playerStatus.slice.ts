import { createSlice } from '@reduxjs/toolkit'

export type playerStatusState = {
  startedTyping: boolean
  finishedTyping: boolean
}

const initialState: playerStatusState = {
  startedTyping: false,
  finishedTyping: false,
}

export const playerStatusSlice = createSlice({
  name: 'playerStatus',
  initialState,
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
