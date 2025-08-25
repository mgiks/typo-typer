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
                  plugins: {
                    legend: { display: false },
                    tooltip: { position: 'nearest' },
                  },
                  scales: {
                    x: {
                      grid: {
                        display: false,
                      },
                      title: {
                        display: true,
                        text: 'Seconds',
                      },
                    },
                    y: {
                      title: {
                        display: true,
                        text: 'WPM',
                      },
                      min: 0,
                      ticks: { stepSize: 1 },
                    },
                    errors: {
                      grid: {
                        display: false,
                      },
                      title: { display: true, text: 'Errors' },
                      axis: 'y',
                      min: 0,
                      position: 'right',
                      ticks: { stepSize: 1 },
                    },
                  },
                }}
                data={{
                  labels: (hasTypedForLessThanASecond
                    ? resultGraphData.filter((point) => point.time < 1000)
                    : resultGraphData.filter((point) =>
                      point.time > 0 && point.time % 1000 == 0
                    ))
                    .map((point) => String(point.time / 1000)).concat(['LRM']),
                  datasets: [{
                    data: (hasTypedForLessThanASecond
                      ? resultGraphData.filter((point) => point.time < 1000)
                      : resultGraphData.filter((point) =>
                        point.time > 0 && point.time % 1000 == 0
                      ))
                      .map((point) => point.wpm).concat([adjustedWpm]),
                  }, {
                    data: resultGraphData.map((point) => point.errs).concat([
                      totalKeysPressed - correctKeysPressed,
                    ]),
                    yAxisID: 'errors',
                  }, {
                    data: resultGraphData.map((point) =>
                      point.acc
                    ).concat([
                      acc,
                    ]),
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
