import { render, screen } from '@testing-library/react'
import Cursor, { type CursorProps } from './Cursor.tsx'

describe('Cursor', async () => {
  it("should be in the document when 'visible' prop is true", () => {
    renderCursor({ visible: true })

    expect(screen.getByRole('marquee')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", async () => {
    renderCursor({ visible: false })

    expect(screen.queryByRole('marquee')).not.toBeInTheDocument()
  })
})

const defaultProps: CursorProps = {
  ref: null,
  visible: false,
}

function renderCursor(overrides: Partial<CursorProps>) {
  return render(<Cursor {...defaultProps} {...overrides} />)
}
