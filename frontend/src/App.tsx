import './App.css'
import TypingPage from './components/TypingPage'

export const API_URL = import.meta.env.VITE_BACKEND_URL ??
  'http://localhost:8080'

function App() {
  return (
    <div>
      <TypingPage />
    </div>
  )
}

export default App
