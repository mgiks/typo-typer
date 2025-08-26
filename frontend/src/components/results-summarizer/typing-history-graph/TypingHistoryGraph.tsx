import { Line } from 'react-chartjs-2'
import type {
  LastRecordedMoment,
  TypingHistoryPoint,
} from '../../../slices/typingHistory.slice'

type TypingHistoryGraphProps = {
  timedData: TypingHistoryPoint[]
  lastRecordedMoment: LastRecordedMoment
}

function TypingHistoryGraph(
  { timedData, lastRecordedMoment }: TypingHistoryGraphProps,
) {
  const latestTime = timedData.at(-1)?.timeInSeconds ?? 0
  const showDataBelowOneSecond = 1 >= latestTime

  return (
    <div className='results-summarizer__graph-wrapper'>
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
              grid: { display: false },
              title: { display: true, text: 'Seconds' },
            },
            y: {
              title: { display: true, text: 'WPM' },
              min: 0,
              ticks: { stepSize: 1 },
            },
            errors: {
              grid: { display: false },
              title: { display: true, text: 'Errors' },
              axis: 'y',
              min: 0,
              position: 'right',
              ticks: { stepSize: 1 },
            },
          },
        }}
        data={{
          labels: filterData(timedData, showDataBelowOneSecond).map((point) =>
            String(point.timeInSeconds)
          ).concat(['LRM']),
          datasets: [
            {
              data: filterData(timedData, showDataBelowOneSecond).map((point) =>
                point.wpm
              ).concat([lastRecordedMoment.wpm]),
            },
            {
              data: filterData(timedData, showDataBelowOneSecond).map((point) =>
                point.errs
              ).concat([lastRecordedMoment.errs]),
              yAxisID: 'errors',
            },
          ],
        }}
      />
    </div>
  )
}

function filterData(
  data: TypingHistoryPoint[],
  showDataBelowOneSecond: boolean,
) {
  return showDataBelowOneSecond
    ? data.filter((point) => point.timeInSeconds < 1)
    : data.filter((point) =>
      point.timeInSeconds > 0 && point.timeInSeconds % 1 == 0
    )
}

export default TypingHistoryGraph
