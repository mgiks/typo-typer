import { render, screen } from '@testing-library/react'
import ReminderToFocus from './ReminderToFocus'

describe('ReminderToFocus', () => {
  it("should appear when 'focused' prop is false", async () => {
    render(
      <ReminderToFocus
        focused={false}
      />,
    )

    expect(await screen.findByTestId('reminder-to-focus')).toBeInTheDocument()
  })

  it("should be hidden when 'focused' prop is true", async () => {
    render(
      <ReminderToFocus
        focused={true}
      />,
    )

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()
  })
})
