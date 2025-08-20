import { createSlice } from '@reduxjs/toolkit'

export type isUserTypingState = {
  value: boolean
}

const initialState: isUserTypingState = {
  value: false,
}

export const isUserTypingSlice = createSlice({
  name: 'isUserTyping',
  initialState,
  reducers: {
    userStartedTyping: (state) => {
      state.value = true
    },
    userStoppedTyping: (state) => {
      state.value = false
    },
  },
})

export const { userStartedTyping, userStoppedTyping } =
  isUserTypingSlice.actions

export default isUserTypingSlice.reducer
