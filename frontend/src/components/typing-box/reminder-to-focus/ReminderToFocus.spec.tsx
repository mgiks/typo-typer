import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import ReminderToFocus from './ReminderToFocus'
import { useState } from 'react'

describe('ReminderToFocus', () => {
  it("should appear when 'focused' prop is false", async () => {
    render(<ReminderToFocus focused={false} focusSetter={() => null} />)

    expect(await screen.findByTestId('reminder-to-focus')).toBeInTheDocument()
  })

  it("should be hidden when 'focused' prop is true", async () => {
    render(<ReminderToFocus focused={true} focusSetter={() => null} />)

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()
  })

  it('should hide on click', async () => {
    function Wrapper() {
      const [focused, setFocused] = useState(false)
      return <ReminderToFocus focused={focused} focusSetter={setFocused} />
    }

    render(<Wrapper />)
    const reminderToFocus = await screen.findByTestId('reminder-to-focus')

    expect(reminderToFocus).toBeInTheDocument()

    const user = userEvent.setup()

    await user.click(reminderToFocus)

    expect(reminderToFocus).not.toBeInTheDocument()
  })
})
