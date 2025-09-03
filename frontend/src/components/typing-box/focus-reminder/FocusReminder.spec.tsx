import { act, render, screen } from '@testing-library/react'
import FocusReminder from './FocusReminder'

describe('FocusReminder', async () => {
  beforeEach(() => vi.useFakeTimers())

  afterEach(() => vi.useRealTimers())

  it("should be in the document after some time when 'visible' prop is true", async () => {
    render(<FocusReminder show={true} />)

    act(() => (vi.runAllTimers(), vi.useRealTimers()))

    expect(await screen.findByRole('status')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", () => {
    render(<FocusReminder show={false} />)

    act(() => vi.runAllTimers())

    expect(screen.queryByRole('status')).not.toBeInTheDocument()
  })
})
