import 'chart.js/auto'
import { Line } from 'react-chartjs-2'
import './ResultsSummarizer.scss'
import { useAppSelector } from '../../hooks'
import {
  calculateAccuracy,
  calculateAdjustedWpm,
  calculateRawWpm,
} from '../../utils/typing-results.utils'
type ResultsSummarizerProps = { forceNoChart?: boolean }

function ResultsSummarizer({ forceNoChart }: ResultsSummarizerProps) {
  const playerFinishedTyping = useAppSelector((state) =>
    state.playerStatus.finishedTyping
  )
  const timeElapsedInMinutes = useAppSelector((state) =>
    state.typingStats.timeElapsedInMinutes
  )
  const totalKeysPressed = useAppSelector((state) =>
    state.typingStats.totalKeysPressed
  )
  const correctKeysPressed = useAppSelector((state) =>
    state.typingStats.correctKeysPressed
  )
  const resultGraphData = useAppSelector((state) => state.resultGraph.data)

  function buildResults() {
    const hasTypedForLessThanASecond = 1 > timeElapsedInMinutes * 60
    const rawWpm = calculateRawWpm(timeElapsedInMinutes, totalKeysPressed)
    const acc = calculateAccuracy(totalKeysPressed, correctKeysPressed)
    const adjustedWpm = calculateAdjustedWpm(rawWpm, acc)

    return (
      <div aria-label='Results summary' className='results-summarizer'>
        <div className='results-summarizer__wpm-acc-section'>
          <div className='results-summarizer__wpm'>
            <div>wpm</div> <div>{adjustedWpm}</div>
          </div>
          <div className='results-summarizer__acc'>
            <div>acc</div>
            <div>{acc * 100}%</div>
          </div>
        </div>
        {forceNoChart
          ? null
          : (
            <div className='results-summarizer__chart-wrapper'>
              <Line
                options={{
                  maintainAspectRatio: false,
                  responsive: true,
                  plugins: { legend: { display: false } },
                  scales: {
                    y: { min: 0, ticks: { stepSize: 1 } },
                    x: { min: 0 },
                  },
                }}
                data={{
                  labels: (hasTypedForLessThanASecond
                    ? resultGraphData.filter((point) => point.time < 1000)
                    : resultGraphData.filter((point) => point.time >= 1000))
                    .map((point) => String(point.time)).concat(['LRM']),
                  datasets: [{
                    data: (hasTypedForLessThanASecond
                      ? resultGraphData.filter((point) => point.time < 1000)
                      : resultGraphData.filter((point) => point.time >= 1000))
                      .map((point) => point.wpm).concat([adjustedWpm]),
                  }],
                }}
              />
            </div>
          )}
      </div>
    )
  }

  return playerFinishedTyping ? buildResults() : null
}

export default ResultsSummarizer
