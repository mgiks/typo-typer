import { render, screen } from '@testing-library/react'
import Cursor from './Cursor.tsx'

describe('Cursor', async () => {
  it("should be in the document when 'visible' prop is true", async () => {
    render(<Cursor visible={true} />)

    expect(await screen.findByTestId('cursor')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is false", async () => {
    render(<Cursor visible={false} />)

    expect(screen.queryByTestId('cursor')).not.toBeInTheDocument()
  })
})
