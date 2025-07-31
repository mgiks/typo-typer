import { render, screen } from '@testing-library/react'
import Cursor from './Cursor.tsx'

describe('Cursor', async () => {
  it('should be in the document', async () => {
    render(<Cursor />)

    expect(await screen.findByTestId('cursor')).toBeInTheDocument()
  })
})
