import { render, screen } from '@testing-library/react'
import InputCatcher from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be in the document', async () => {
    render(<InputCatcher />)

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should be initially focused', async () => {
    render(<InputCatcher />)

    expect(await screen.findByTestId('input-catcher')).toHaveFocus()
  })
})
