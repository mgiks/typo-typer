import { Line } from 'react-chartjs-2'
import { Result } from '../../stores/ResultStore'
import { secondsToStringArray } from './utils/secondsToStringArray'
import './Chart.css'
import {
  CategoryScale,
  Chart as ChartJS,
  LinearScale,
  LineController,
  LineElement,
  PointElement,
  Tooltip,
} from 'chart.js'

ChartJS.register(
  LineController,
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Tooltip,
)

function Chart(
  { finalResults, finalTime }: { finalResults: Result[]; finalTime: number },
) {
  return (
    <div id='chart'>
      <Line
        data={{
          datasets: [
            {
              label: 'WPM',
              data: finalResults.map((result) => result.NWPM),
              borderWidth: 1,
            },
            {
              label: 'Errors',
              data: finalResults.map((
                result,
              ) => (result.errors)),
              yAxisID: 'y1',
            },
          ],
        }}
        options={{
          maintainAspectRatio: false,
          scales: {
            y: {
              title: {
                color: 'black',
                text: 'Words Per Minute',
                align: 'center',
                display: true,
                padding: {
                  top: 30,
                },
              },
              position: 'left',
              min: 1,
            },
            y1: {
              title: {
                color: 'black',
                text: 'Errors',
                align: 'center',
                display: true,
              },
              position: 'right',
              beginAtZero: true,
              ticks: {
                stepSize: 1,
              },
            },
            x: {
              labels: secondsToStringArray(finalTime),
            },
          },
          plugins: {
            tooltip: {
              enabled: true,
            },
          },
        }}
      />
    </div>
  )
}

export function secondsToStringArray(seconds: number) {
  if (seconds <= 0) return []
  return Array(seconds).fill(0).map((_, i) => (seconds - i).toString())
    .reverse()
}

export default Chart
