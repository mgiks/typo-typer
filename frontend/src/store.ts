import { combineReducers, configureStore } from '@reduxjs/toolkit'
import isUserTypingReducer from './slices/isUserTyping.slice'

const rootReducer = combineReducers({ isUserTyping: isUserTypingReducer })

export const store = configureStore({ reducer: rootReducer })

export const setupStore = (preloadedState?: Partial<RootState>) =>
  configureStore({ reducer: rootReducer, preloadedState })

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type AppStore = typeof store
