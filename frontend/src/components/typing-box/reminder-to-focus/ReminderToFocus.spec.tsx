import { render, screen } from '@testing-library/react'
import ReminderToFocus from './ReminderToFocus'

describe('ReminderToFocus', () => {
  it("should be in the document when 'visible' prop is true", async () => {
    render(<ReminderToFocus visible={true} />)

    expect(await screen.findByTestId('reminder-to-focus')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", async () => {
    render(<ReminderToFocus visible={false} />)

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()
  })
})
