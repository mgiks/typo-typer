import { combineReducers, configureStore } from '@reduxjs/toolkit'
import typingDataReducer from './slices/typingData.slice'
import playerStatusReducer from './slices/playerStatus.slice'
import typingHistoryReducer from './slices/typingHistory.slice'
import textDataReducer from './slices/textData.slice'

const rootReducer = combineReducers({
  typingData: typingDataReducer,
  typingHistory: typingHistoryReducer,
  playerStatus: playerStatusReducer,
  textData: textDataReducer,
})

export const store = configureStore({ reducer: rootReducer })

export const setupStore = (preloadedState?: Partial<RootState>) =>
  configureStore({ reducer: rootReducer, preloadedState })

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type AppStore = typeof store
