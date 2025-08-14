import { render, screen } from '@testing-library/react'
import StopWatch from './StopWatch'

describe('StopWatch', async () => {
  it("should be in the document when 'visible' prop is true", async () => {
    render(<StopWatch visible={true} />)

    expect(await screen.findByTestId('timer')).toBeInTheDocument()
  })

  it("should not be in the document when 'visible' prop is true", async () => {
    render(<StopWatch visible={false} />)

    expect(screen.queryByTestId('timer')).not.toBeInTheDocument()
  })
})
