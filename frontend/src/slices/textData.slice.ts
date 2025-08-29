import {
  createAsyncThunk,
  createSlice,
  type PayloadAction,
} from '@reduxjs/toolkit'

export type TextDataState = {
  text: string
  lastTypedIndex: number
  incorrectTextStartIndex: number
}

export type TextResponse = { text: string }

export const TEXTS_URL = 'http://localhost:8000/texts'

export const fetchText = createAsyncThunk(
  'textData/fetchText',
  async () => {
    const text = await fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as TextResponse)
      .then((json) => json.text)
      .catch((err) => err)

    return text
  },
)

export const textDataInitialState: TextDataState = {
  text: '',
  lastTypedIndex: -1,
  incorrectTextStartIndex: -1,
}

const textDateSlice = createSlice({
  name: 'textData',
  initialState: textDataInitialState,
  reducers: {
    setTextTo: (state, action: PayloadAction<string>) => {
      state.text = action.payload
    },
    setLastTypedIndexTo: (state, action: PayloadAction<number>) => {
      state.lastTypedIndex = action.payload
    },
    setIncorrectTextStartIndexTo: (state, action: PayloadAction<number>) => {
      state.incorrectTextStartIndex = action.payload
    },
    resetTextData: (state) => {
      state.text = textDataInitialState.text
      state.lastTypedIndex = textDataInitialState.lastTypedIndex
      state.incorrectTextStartIndex =
        textDataInitialState.incorrectTextStartIndex
    },
  },
  extraReducers: (builder) => {
    builder.addCase(fetchText.fulfilled, (state, action) => {
      state.text = action.payload
    })
  },
})

export const {
  setTextTo,
  setLastTypedIndexTo,
  setIncorrectTextStartIndexTo,
  resetTextData,
} = textDateSlice.actions

export default textDateSlice.reducer
