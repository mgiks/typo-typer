import './App.scss'
import ModeSwitcher from './components/mode-switcher/ModeSwitcher'
import ResetButton from './components/reset-button/ResetButton'
import ResultsSummarizer from './components/results-summarizer/ResultsSummarizer'
import StopWatch from './components/stop-watch/StopWatch'
import TypingBox from './components/typing-box/TypingBox'

function App() {
  return (
    <div className='app'>
      <div>
        Header placeholder
      </div>
      <div className='app__middle-section'>
        <div className='app__typing-section'>
          <ModeSwitcher />
          <TypingBox />
          <StopWatch />
        </div>
        <div className='app__result-section'>
          <ResultsSummarizer />
          <ResetButton />
        </div>
      </div>
      <div>
        Footer placeholder
      </div>
    </div>
  )
}

export default App
