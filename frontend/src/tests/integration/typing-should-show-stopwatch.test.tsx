import { screen } from '@testing-library/react'
import StopWatch from '../../components/stop-watch/StopWatch'
import userEvent from '@testing-library/user-event'
import TypingBox from '../../components/typing-box/TypingBox'
import { renderWithProviders } from '../utils'

it('typing should show stopwatch', async () => {
  const user = userEvent.setup()
  renderWithProviders(
    <>
      <TypingBox />
      <StopWatch />
    </>,
  )

  expect(screen.queryByRole('timer')).not.toBeInTheDocument()

  await user.keyboard('t')

  expect(screen.getByRole('timer')).toBeInTheDocument()
})
