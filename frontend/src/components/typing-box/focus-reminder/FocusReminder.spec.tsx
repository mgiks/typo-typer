import { render, screen } from '@testing-library/react'
import FocusReminder from './FocusReminder'

describe('FocusReminder', () => {
  it("should be in the document when 'visible' prop is true", () => {
    render(<FocusReminder visible={true} />)

    expect(screen.getByRole('status')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", () => {
    render(<FocusReminder visible={false} />)

    expect(screen.queryByRole('status')).not.toBeInTheDocument()
  })
})
