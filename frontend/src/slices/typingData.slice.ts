import { createSlice } from '@reduxjs/toolkit'

export type TypingDataState = {
  totalKeysPressed: number
  correctKeysPressed: number
}

export const typingDataInitialState: TypingDataState = {
  totalKeysPressed: 0,
  correctKeysPressed: 0,
}

export const typingDataSlice = createSlice({
  name: 'typingData',
  initialState: typingDataInitialState,
  reducers: {
    increaseTotalKeysPressed: (state) => {
      state.totalKeysPressed += 1
    },
    increaseCorrectKeysPressed: (state) => {
      state.correctKeysPressed += 1
    },
    resetTypingData: (state) => {
      state.totalKeysPressed = typingDataInitialState.totalKeysPressed
      state.correctKeysPressed = typingDataInitialState.correctKeysPressed
    },
  },
})

export const {
  increaseTotalKeysPressed,
  increaseCorrectKeysPressed,
  resetTypingData,
} = typingDataSlice.actions

export default typingDataSlice.reducer
