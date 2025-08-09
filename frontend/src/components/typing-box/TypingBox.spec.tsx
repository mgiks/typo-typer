import { act, render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox from './TypingBox.tsx'
import { FOCUS_REMINDER_TIMEOUT_MS } from './input-catcher/InputCatcher.tsx'

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

  it('should include container for text', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should include input catcher', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should show focus reminder when input catcher is blurred', async () => {
    const { findByTestId } = render(<TypingBox />)
    const inputCatcher = await findByTestId('input-catcher')

    expect(screen.queryByTestId('focus-reminder')).not.toBeInTheDocument()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(await findByTestId('focus-reminder')).toBeInTheDocument()
  })

  it('should focus input catcher on click', async () => {
    const { findByTestId } = render(<TypingBox />)
    const typingBox = await screen.findByTestId('typing-box')
    const inputCatcher = await findByTestId('input-catcher')

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    const user = userEvent.setup()
    await user.click(typingBox)

    expect(inputCatcher).toHaveFocus()
  })

  it('should focus input catcher on keypress', async () => {
    render(<TypingBox />)
    const inputCatcher = await screen.findByTestId('input-catcher')
    const user = userEvent.setup()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(inputCatcher).not.toHaveFocus()

    await user.keyboard('o')

    expect(inputCatcher).toHaveFocus()
  })

  it('should hide focus reminder on click', async () => {
    const { findByTestId } = render(<TypingBox />)
    const typingBox = await screen.findByTestId('typing-box')
    const inputCatcher = await findByTestId('input-catcher')

    expect(screen.queryByTestId('focus-reminder')).not.toBeInTheDocument()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(await findByTestId('focus-reminder')).toBeInTheDocument()

    const user = userEvent.setup()
    await user.click(typingBox)

    expect(screen.queryByTestId('focus-reminder')).not.toBeInTheDocument()
  })
})
