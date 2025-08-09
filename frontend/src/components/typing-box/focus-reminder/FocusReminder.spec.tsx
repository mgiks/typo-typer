import { render, screen } from '@testing-library/react'
import FocusReminder from './FocusReminder'

describe('FocusReminder', () => {
  it("should be in the document when 'visible' prop is true", async () => {
    render(<FocusReminder visible={true} />)

    expect(await screen.findByTestId('focus-reminder')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", async () => {
    render(<FocusReminder visible={false} />)

    expect(screen.queryByTestId('focus-reminder')).not.toBeInTheDocument()
  })
})
