import { render, screen } from '@testing-library/react'
import StopWatch from '../../components/stop-watch/StopWatch'
import userEvent from '@testing-library/user-event'
import { Provider } from 'react-redux'
import { store } from '../../store'
import TypingBox from '../../components/typing-box/TypingBox'

it('typing should show stopwatch', async () => {
  const user = userEvent.setup()
  render(
    <Provider store={store}>
      <TypingBox />
      <StopWatch />
    </Provider>,
  )

  expect(screen.queryByRole('timer')).not.toBeInTheDocument()

  await user.keyboard('t')

  expect(screen.getByRole('timer')).toBeInTheDocument()
})
