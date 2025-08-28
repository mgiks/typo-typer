import userEvent from '@testing-library/user-event'
import ResultsSummarizer from '../../components/results-summarizer/ResultsSummarizer'
import StopWatch from '../../components/stop-watch/StopWatch'
import TypingBox from '../../components/typing-box/TypingBox'
import { TEXT_FIXTURE } from '../fixtures'
import { renderWithProviders } from '../utils'
import { screen } from '@testing-library/react'

it('typing results should make sense', async () => {
  const user = userEvent.setup()
  renderWithProviders(
    <>
      <TypingBox forcedText={TEXT_FIXTURE} />
      <StopWatch />
      <ResultsSummarizer forceNoGraph={true} />
    </>,
  )

  await user.keyboard('Dest text.')

  expect(screen.getByLabelText('Results summary')).toBeInTheDocument()
  expect(screen.getByText('0')).toBeInTheDocument()
  expect(screen.getByText('0%')).toBeInTheDocument()
})
