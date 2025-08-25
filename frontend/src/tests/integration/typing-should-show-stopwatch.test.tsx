import { screen } from '@testing-library/react'
import StopWatch from '../../components/stop-watch/StopWatch'
import userEvent from '@testing-library/user-event'
import TypingBox from '../../components/typing-box/TypingBox'
import { renderWithProviders } from '../utils'
import { TEXT_FIXTURE } from '../fixtures'

it('typing should show stopwatch', async () => {
  const user = userEvent.setup()
  renderWithProviders(
    <>
      <TypingBox initialText={TEXT_FIXTURE} />
      <StopWatch />
    </>,
  )

  expect(screen.queryByRole('timer')).not.toBeInTheDocument()

  await user.keyboard('t')

  expect(screen.getByRole('timer')).toBeInTheDocument()
})
