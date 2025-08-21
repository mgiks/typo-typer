import { createSlice, type PayloadAction } from '@reduxjs/toolkit'

export type typingStatsState = {
  totalKeysPressed: number
  correctKeysPressed: number
  timeElapsedInMinutes: number
}

export const typingStatsInitialState: typingStatsState = {
  totalKeysPressed: 0,
  correctKeysPressed: 0,
  timeElapsedInMinutes: 0,
}

export const typingStatsSlice = createSlice({
  name: 'typingStats',
  initialState: typingStatsInitialState,
  reducers: {
    increaseTotalKeysPressed: (state) => {
      state.totalKeysPressed += 1
    },
    increaseCorrectKeysPressed: (state) => {
      state.correctKeysPressed += 1
    },
    setTimeElapsedInMinutesTo: (state, action: PayloadAction<number>) => {
      state.timeElapsedInMinutes = action.payload
    },
  },
})

export const {
  increaseTotalKeysPressed,
  increaseCorrectKeysPressed,
  setTimeElapsedInMinutesTo,
} = typingStatsSlice.actions

export default typingStatsSlice.reducer
