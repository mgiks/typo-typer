import { act, screen } from '@testing-library/react'
import StopWatch from './StopWatch'
import { renderWithProviders } from '../../tests/utils'

describe('StopWatch', async () => {
  it('should update every second', async () => {
    vi.useFakeTimers()

    renderWithProviders(<StopWatch forceVisible={true} />)

    expect(screen.getByText('0')).toBeInTheDocument()

    act(() => vi.advanceTimersByTime(1000))

    expect(screen.getByText('1')).toBeInTheDocument()

    act(() => vi.advanceTimersByTime(1000))

    expect(screen.getByText('2')).toBeInTheDocument()

    vi.useRealTimers()
  })
})
