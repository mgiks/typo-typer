import { configureStore } from '@reduxjs/toolkit'
import isUserTypingReducer from './slices/isUserTyping.slice'

export const store = configureStore({
  reducer: { isUserTyping: isUserTypingReducer },
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type AppStore = typeof store
