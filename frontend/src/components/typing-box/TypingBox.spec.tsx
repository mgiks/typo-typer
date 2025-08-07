import { act, render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox from './TypingBox.tsx'
import { BLUR_TIMEOUT } from './input-catcher/InputCatcher.tsx'

const handlers = [
  http.get('http://localhost:8000/texts', () => {
    return HttpResponse.json({
      text: 'Some text',
    })
  }),
]

const server = setupServer(...handlers)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('TypingBox', async () => {
  it('should be in the document', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('typing-box')).toBeInTheDocument()
  })

  it('should display fetched text on initial render', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('typing-box')).toHaveTextContent(
      'Some text',
    )
  })

  it('should include text-container', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should include input-catcher', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should include reminder-to-focus', async () => {
    const { findByTestId } = render(<TypingBox />)
    const inputCatcher = await findByTestId('input-catcher')

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()

    vi.useFakeTimers()
    inputCatcher.blur()
    act(() => vi.advanceTimersByTime(BLUR_TIMEOUT))
    vi.useRealTimers()

    expect(await screen.findByTestId('reminder-to-focus')).toBeInTheDocument()
  })

  it('should focus input-catcher on click', async () => {
    const { findByTestId } = render(<TypingBox />)
    const typingBox = await screen.findByTestId('typing-box')
    const inputCatcher = await findByTestId('input-catcher')

    inputCatcher.blur()

    const user = userEvent.setup()
    await user.click(typingBox)

    expect(inputCatcher).toHaveFocus()
  })

  it('should hide reminder-to-focus on click', async () => {
    const { findByTestId } = render(<TypingBox />)
    const typingBox = await screen.findByTestId('typing-box')
    const inputCatcher = await findByTestId('input-catcher')

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()

    vi.useFakeTimers()
    inputCatcher.blur()
    act(() => vi.advanceTimersByTime(BLUR_TIMEOUT))
    vi.useRealTimers()

    expect(await findByTestId('reminder-to-focus')).toBeInTheDocument()

    const user = userEvent.setup()
    await user.click(typingBox)

    expect(screen.queryByTestId('reminder-to-focus')).not.toBeInTheDocument()
  })
})
