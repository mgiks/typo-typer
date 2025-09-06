import { createSlice, type PayloadAction } from '@reduxjs/toolkit'

export type MultiplayerDataState = {
  name: string
  averageWpm: number
}

export const multiplayerDataInitialState: MultiplayerDataState = {
  name: 'placeholder bandit',
  // World Average WPM (approximately)
  averageWpm: 40,
}

export const multiplayerDataSlice = createSlice({
  name: 'multiplayerData',
  initialState: multiplayerDataInitialState,
  reducers: {
    setName: (state, action: PayloadAction<string>) => {
      state.name = action.payload
    },
    setAverageWpm: (state, action: PayloadAction<number>) => {
      state.averageWpm = action.payload
    },
  },
})

export const { setName, setAverageWpm } = multiplayerDataSlice.actions

export default multiplayerDataSlice.reducer
