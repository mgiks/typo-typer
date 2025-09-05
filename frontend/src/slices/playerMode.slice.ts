import { createSlice } from '@reduxjs/toolkit'

export type PlayerModeState = {
  mode: 'singlePlayer' | 'multiPlayer'
}

export const playerModeInitialState: PlayerModeState = {
  mode: 'singlePlayer',
}

export const playerModeSlice = createSlice({
  name: 'playerMode',
  initialState: playerModeInitialState,
  reducers: {
    setModeToSinglePlayer: (state) => {
      state.mode = 'singlePlayer'
    },
    setModeToMultiPlayer: (state) => {
      state.mode = 'multiPlayer'
    },
  },
})

export const { setModeToSinglePlayer, setModeToMultiPlayer } =
  playerModeSlice.actions

export default playerModeSlice.reducer
