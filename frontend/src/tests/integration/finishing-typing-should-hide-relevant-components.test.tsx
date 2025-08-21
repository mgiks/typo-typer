import { screen } from '@testing-library/react'
import StopWatch from '../../components/stop-watch/StopWatch'
import userEvent from '@testing-library/user-event'
import TypingBox from '../../components/typing-box/TypingBox'
import { renderWithProviders } from '../utils'
import { TEXT_FIXTURE } from '../fixtures'

it('finishing typing should hide relevant components', async () => {
  const user = userEvent.setup()
  renderWithProviders(
    <>
      <TypingBox initialText={TEXT_FIXTURE} />
      <StopWatch />
    </>,
  )

  await user.keyboard('T')

  expect(screen.getByRole('region')).toBeInTheDocument()
  expect(screen.getByRole('timer')).toBeInTheDocument()

  await user.keyboard('est text.')

  expect(screen.queryByRole('region')).not.toBeInTheDocument()
  expect(screen.queryByRole('timer')).not.toBeInTheDocument()
})
