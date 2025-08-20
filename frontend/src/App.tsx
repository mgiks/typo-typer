import './App.scss'
import StopWatch from './components/stop-watch/StopWatch'
import TypingBox from './components/typing-box/TypingBox'

function App() {
  return (
    <div className='app'>
      <div />
      <div className='app__typing-section'>
        <TypingBox />
        <StopWatch />
      </div>
      <div />
    </div>
  )
}

export default App
