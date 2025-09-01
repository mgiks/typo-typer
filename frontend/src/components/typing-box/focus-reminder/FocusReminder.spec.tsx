import { act, render, screen } from '@testing-library/react'
import FocusReminder, { FOCUS_REMINDER_TIMEOUT_MS } from './FocusReminder'

describe('FocusReminder', async () => {
  it("should be in the document after a timeout when 'visible' prop is true", () => {
    vi.useFakeTimers()
    render(<FocusReminder show={true} />)

    act(() => {
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(screen.getByRole('status')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", () => {
    render(<FocusReminder show={false} />)

    expect(screen.queryByRole('status')).not.toBeInTheDocument()
  })
})
