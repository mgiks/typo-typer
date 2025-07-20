import { render, screen } from '@testing-library/react'
import TextContainer from './TextContainer.tsx'

describe('TextContainer', async () => {
  it('should be in the document', async () => {
    render(<TextContainer text='' />)

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should display passed text', async () => {
    const text = 'Test text.'

    render(<TextContainer text={text} />)

    expect(await screen.findByText(text)).toBeInTheDocument()
  })

  it('should include cursor', async () => {
    const { findByTestId } = render(<TextContainer text='' />)

    expect(await findByTestId('cursor')).toBeInTheDocument()
  })
})
