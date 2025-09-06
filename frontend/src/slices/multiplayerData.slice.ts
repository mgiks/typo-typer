import { createSlice, type PayloadAction } from '@reduxjs/toolkit'

export type MultiplayerDataState = {
  name: string
}

export const multiplayerDataInitialState: MultiplayerDataState = {
  name: 'placeholder bandit',
}

export const multiplayerDataSlice = createSlice({
  name: 'multiplayerData',
  initialState: multiplayerDataInitialState,
  reducers: {
    setName: (state, action: PayloadAction<string>) => {
      state.name = action.payload
    },
  },
})

export const { setName } = multiplayerDataSlice.actions

export default multiplayerDataSlice.reducer
