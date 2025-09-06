import './App.scss'
import ModeSwitcher from './components/mode-switcher/ModeSwitcher'
import ProfileCorner from './components/profile-corner/ProfileCorner'
import ResetButton from './components/reset-button/ResetButton'
import ResultsSummarizer from './components/results-summarizer/ResultsSummarizer'
import StopWatch from './components/stop-watch/StopWatch'
import TypingBox from './components/typing-box/TypingBox'

function App() {
  return (
    <div className='app'>
      <div className='app__header'>
        Header placeholder
        <ProfileCorner />
      </div>
      <div className='app__middle-section'>
        <div className='app__upper-typing-section'>
          <ModeSwitcher />
          <TypingBox />
        </div>
        <div className='app__lower-typing-section'>
          <ResetButton />
          <StopWatch />
        </div>
        <div className='app__result-section'>
          <ResultsSummarizer />
        </div>
      </div>
      <div>
        Footer placeholder
      </div>
    </div>
  )
}

export default App
