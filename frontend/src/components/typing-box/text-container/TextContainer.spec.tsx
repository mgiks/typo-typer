import { render, screen } from '@testing-library/react'
import TextContainer from './TextContainer.tsx'

describe('TextContainer', async () => {
  it('should be in the document', async () => {
    render(<TextContainer text='' />)

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should display passed text', async () => {
    render(<TextContainer text='Test text.' />)

    expect(await screen.findByText('Test text.')).toBeInTheDocument()
  })
})
