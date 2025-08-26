import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

export type TypingHistoryPoint = {
  timeInSeconds: number
  wpm: number
  acc: number
  errs: number
}

export type LastRecordedMoment = Omit<TypingHistoryPoint, 'timeInSeconds'>
export type TypingHistoryState = {
  timedData: TypingHistoryPoint[]
  lastRecordedMoment: LastRecordedMoment
}

export const TypingHistoryInitialState: TypingHistoryState = {
  timedData: [],
  lastRecordedMoment: { wpm: 0, acc: 0, errs: 0 },
}

export const typingHistorySlice = createSlice({
  name: 'TimedTyping',
  initialState: TypingHistoryInitialState,
  reducers: {
    setTypingHistory: (
      state,
      action: PayloadAction<TypingHistoryPoint[]>,
    ) => {
      state.timedData = action.payload
    },
    addTypingHistoryPoint: (
      state,
      action: PayloadAction<TypingHistoryPoint>,
    ) => {
      state.timedData.push({
        timeInSeconds: action.payload.timeInSeconds,
        wpm: action.payload.wpm,
        acc: action.payload.acc,
        errs: action.payload.errs,
      })
    },
    setLastRecordedMoment: (
      state,
      action: PayloadAction<LastRecordedMoment>,
    ) => {
      state.lastRecordedMoment = action.payload
    },
  },
})

export const {
  setTypingHistory,
  addTypingHistoryPoint,
  setLastRecordedMoment,
} = typingHistorySlice.actions

export default typingHistorySlice.reducer
