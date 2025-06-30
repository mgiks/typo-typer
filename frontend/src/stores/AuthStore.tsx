import { create } from 'zustand'

type AuthActions = {
  setAccessToken: (token: string) => void
}

type AuthState = {
  accessToken: string
  actions: AuthActions
}

const useAuthStore = create<AuthState>()((set) => ({
  accessToken: '',
  actions: {
    setAccessToken: (token: string) => set({ accessToken: token }),
  },
}))

export const useAccessToken = () => useAuthStore((state) => state.accessToken)
export const useAuthActions = () => useAuthStore((state) => state.actions)
