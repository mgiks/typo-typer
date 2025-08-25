import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'

export type ResultGraphPointData = {
  time: number
  wpm: number
  acc: number
  errs: number
}

export type ResultGraphState = {
  data: ResultGraphPointData[]
}

export const ResultGraphInitialState: ResultGraphState = {
  data: [],
}

export const resultGraphSlice = createSlice({
  name: 'resultGraph',
  initialState: ResultGraphInitialState,
  reducers: {
    setResultGraphData: (
      state,
      action: PayloadAction<ResultGraphPointData[]>,
    ) => {
      state.data = action.payload
    },
    addResultGraphPoint: (
      state,
      action: PayloadAction<ResultGraphPointData>,
    ) => {
      state.data.push({
        time: action.payload.time,
        wpm: action.payload.wpm,
        acc: action.payload.acc,
        errs: action.payload.errs,
      })
    },
  },
})

export const { setResultGraphData, addResultGraphPoint } =
  resultGraphSlice.actions

export default resultGraphSlice.reducer
