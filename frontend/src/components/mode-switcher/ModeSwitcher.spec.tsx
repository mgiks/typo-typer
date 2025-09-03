import { render, screen } from '@testing-library/react'
import ModeSwitcher from './ModeSwitcher'
import userEvent from '@testing-library/user-event'

describe('ModeSwitcher', async () => {
  it('should show modal on click', async () => {
    const user = userEvent.setup()
    render(
      <>
        <ModeSwitcher />
        <div id='modal-container' />
      </>,
    )

    await user.click(screen.getByRole('button'))

    expect(screen.findByText(/^choose mode$/i))
  })
})
