import './InputCatcher.scss'
import { render, screen } from '@testing-library/react'
import InputCatcher from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be initially focused', async () => {
    render(<InputCatcher />)

    expect(await screen.findByTestId('input-catcher')).toHaveFocus()
  })
})
