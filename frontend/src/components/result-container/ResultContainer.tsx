import './ResultContainer.css'
import { useCursorIndex, useTextRefreshCount } from '../../stores/TextStore'
import { useEffect } from 'react'
import {
  useCorrectKeyCount,
  useIsDoneTyping,
  useTypingStatsActions,
  useTypingTime,
  useWrongKeyCount,
} from '../../stores/TypingStatsStore'
import {
  Result,
  useResultActions,
  useResultsPerSecond,
} from '../../stores/ResultStore'
import Chart from './Chart'
import { calculateTypingAccuracyAndWPM } from '../../utils/utils'

function ResultContainer() {
  const isDoneTyping = useIsDoneTyping()
  const typingTime = useTypingTime()
  const errors = useWrongKeyCount()
  const correctKeyPresses = useCorrectKeyCount()
  const textRefreshCount = useTextRefreshCount()
  const results = useResultsPerSecond()
  const cursorIndex = useCursorIndex()

  const { clearResults } = useResultActions()
  const { startTypingGame } = useTypingStatsActions()

  useEffect(() => {
    clearResults()
    startTypingGame()
  }, [textRefreshCount])

  if (!isDoneTyping) {
    return null
  }

  const { GWPM, NWPM, typingAccuracy } = calculateTypingAccuracyAndWPM(
    cursorIndex,
    typingTime,
    correctKeyPresses,
    errors,
  )

  const result: Result = {
    GWPM: GWPM,
    NWPM: NWPM,
    typingAccuracy: typingAccuracy,
    time: typingTime + 1,
    errors: errors,
  }

  // Needed to prevent contradiction of graph data and final result
  const finalResults = [...results, result]
  const finalTime = typingTime + 1

  return (
    <div
      id='result-container'
      className={isDoneTyping ? undefined : 'invisible'}
    >
      <div id='result-wrapper'>
        <div id='stats'>
          <div id='wpm' className='result-stat'>
            <div className='stat-label'>
              wpm
            </div>
            <div className='stat-value'>
              {NWPM}
            </div>
          </div>
          <div id='accuracy' className='result-stat'>
            <div className='stat-label'>
              accuracy
            </div>
            <div className='stat-value'>
              {typingAccuracy * 100}%
            </div>
          </div>
        </div>
        <div id='additional-stats'>
          <div id='raw-wpm' className='result-stat'>
            <div className='small-stat-label'>
              raw wpm
            </div>
            <div className='small-stat-value'>
              {GWPM}
            </div>
          </div>
          <div id='time' className='result-stat'>
            <div className='small-stat-label'>
              time
            </div>
            <div className='small-stat-value'>
              {typingTime}s
            </div>
          </div>
          <div id='errors' className='result-stat'>
            <div className='small-stat-label'>
              errors
            </div>
            <div className='small-stat-value'>
              {errors}
            </div>
          </div>
        </div>
        <Chart finalResults={finalResults} finalTime={finalTime} />
      </div>
    </div>
  )
}

export default ResultContainer
