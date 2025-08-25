import { combineReducers, configureStore } from '@reduxjs/toolkit'
import playerStatusReducer from './slices/playerStatus.slice'
import typingStatsReducer from './slices/typingStats.slice'
import resultGraphReducer from './slices/resultGraph.slice'

const rootReducer = combineReducers({
  playerStatus: playerStatusReducer,
  typingStats: typingStatsReducer,
  resultGraph: resultGraphReducer,
})

export const store = configureStore({ reducer: rootReducer })

export const setupStore = (preloadedState?: Partial<RootState>) =>
  configureStore({ reducer: rootReducer, preloadedState })

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type AppStore = typeof store
