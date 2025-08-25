import './App.scss'
import ResultsSummarizer from './components/results-summarizer/ResultsSummarizer'
import StopWatch from './components/stop-watch/StopWatch'
import TypingBox from './components/typing-box/TypingBox'

function App() {
  return (
    <div className='app'>
      <div />
      <div className='app__middle-section'>
        <TypingBox initialText='Bro.' />
        <StopWatch />
        <ResultsSummarizer />
      </div>
      <div />
    </div>
  )
}

export default App
