import 'chart.js/auto'
import './ResultsSummarizer.scss'
import { useAppSelector } from '../../hooks'
import TypingHistoryGraph from './typing-history-graph/TypingHistoryGraph'

type ResultsSummarizerProps = { forceNoGraph?: boolean }

function ResultsSummarizer({ forceNoGraph }: ResultsSummarizerProps) {
  const playerFinishedTyping = useAppSelector((state) =>
    state.playerStatus.finishedTyping
  )
  const timedtypingHistory = useAppSelector((state) =>
    state.typingHistory.timedData
  )
  const lastRecordedMoment = useAppSelector((state) =>
    state.typingHistory.lastRecordedMoment
  )
  const latestWpm = lastRecordedMoment.wpm
  const latestAcc = lastRecordedMoment.acc

  function buildResults() {
    return (
      <div aria-label='Results summary' className='results-summarizer'>
        <div className='results-summarizer__wpm-acc-section'>
          <div className='results-summarizer__wpm'>
            <div>wpm</div> <div>{latestWpm}</div>
          </div>
          <div className='results-summarizer__acc'>
            <div>acc</div> <div>{latestAcc * 100}%</div>
          </div>
        </div>
        <div className='results-summarizer__graph-wrapper'>
          {forceNoGraph ? null : (
            <TypingHistoryGraph
              timedData={timedtypingHistory}
              lastRecordedMoment={lastRecordedMoment}
            />
          )}
        </div>
      </div>
    )
  }

  return playerFinishedTyping ? buildResults() : null
}

export default ResultsSummarizer
