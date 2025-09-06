import { createSlice } from '@reduxjs/toolkit'

export type PlayerModeState = {
  mode: 'Solo' | 'Multiplayer'
}

export const playerModeInitialState: PlayerModeState = {
  mode: 'Solo',
}

export const playerModeSlice = createSlice({
  name: 'playerMode',
  initialState: playerModeInitialState,
  reducers: {
    setModeToSinglePlayer: (state) => {
      state.mode = 'Solo'
    },
    setModeToMultiPlayer: (state) => {
      state.mode = 'Multiplayer'
    },
  },
})

export const { setModeToSinglePlayer, setModeToMultiPlayer } =
  playerModeSlice.actions

export default playerModeSlice.reducer
