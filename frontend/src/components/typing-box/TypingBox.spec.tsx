import { render, screen } from '@testing-library/react'
import TypingBox from './TypingBox.tsx'

describe('TypingBox', async () => {
  it('should be in the document', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('typing-box')).toBeInTheDocument()
  })
})
