import { createSlice } from '@reduxjs/toolkit'

export type typingDataState = {
  totalKeysPressed: number
  correctKeysPressed: number
}

export const typingDataInitialState: typingDataState = {
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
  },
})

export const {
  increaseTotalKeysPressed,
  increaseCorrectKeysPressed,
} = typingDataSlice.actions

export default typingDataSlice.reducer
