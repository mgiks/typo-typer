import 'chart.js/auto'
import { Line } from 'react-chartjs-2'
import './ResultsSummarizer.scss'
import { useAppSelector } from '../../hooks'
import type { typingStatsState } from '../../slices/typingStats.slice'

type ResultsSummarizerProps = { forceNoChart?: boolean }

function ResultsSummarizer({ forceNoChart }: ResultsSummarizerProps) {
  const playerFinishedTyping = useAppSelector((state) =>
    state.playerStatus.finishedTyping
  )
  const totalKeysPressed = useAppSelector((state) =>
    state.typingStats.totalKeysPressed
  )
  const correctKeysPressed = useAppSelector((state) =>
    state.typingStats.correctKeysPressed
  )
  const timeElapsedInMinutes = useAppSelector((state) =>
    state.typingStats.timeElapsedInMinutes
  )

  function buildResults(
    { totalKeysPressed, correctKeysPressed, timeElapsedInMinutes }:
      typingStatsState,
  ) {
    const wpm = Math.floor(totalKeysPressed / 5 / timeElapsedInMinutes)
    const acc = correctKeysPressed / totalKeysPressed

    return (
      <div aria-label='Results summary' className='results-summarizer'>
        <div className='results-summarizer__wpm-acc-section'>
          <div className='results-summarizer__wpm'>
            <div>wpm</div> <div>{wpm * acc}</div>
          </div>
          <div className='results-summarizer__acc'>
            <div>acc</div> <div>{acc * 100}%</div>
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
                  labels: Array.from({ length: 100 }, (_, i) => i),
                  datasets: [{
                    data: Array.from({ length: 100 }, (_, i) =>
                      Math.sin(i / 10)),
                  }],
                }}
              />
            </div>
          )}
      </div>
    )
  }

  return playerFinishedTyping
    ? buildResults({
      totalKeysPressed: totalKeysPressed,
      correctKeysPressed: correctKeysPressed,
      timeElapsedInMinutes: timeElapsedInMinutes,
    })
    : null
}

export default ResultsSummarizer
