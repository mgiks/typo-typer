import './App.scss'
import ResetButton from './components/reset-button/ResetButton'
import ResultsSummarizer from './components/results-summarizer/ResultsSummarizer'
import StopWatch from './components/stop-watch/StopWatch'
import TypingBox from './components/typing-box/TypingBox'

function App() {
  return (
    <div className='app'>
      <div />
      <div className='app__middle-section'>
        <TypingBox />
        <StopWatch />
        <div className='app__result-section'>
          <ResultsSummarizer />
          <ResetButton />
        </div>
      </div>
      <div />
    </div>
  )
}

export default App
