import { render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox from './TypingBox.tsx'

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
    render(<TypingBox />)
    render(<div data-testid='focus-stealer' />)

    const focusStealer = await screen.findByTestId('focus-stealer')

    const user = userEvent.setup()
    await user.click(focusStealer)

    expect(await screen.findByTestId('reminder-to-focus')).toBeInTheDocument()
  })

  it('should focus input-catcher on click', async () => {
    const { findByTestId } = render(<TypingBox />)
    const user = userEvent.setup()
    const typingBox = await screen.findByTestId('typing-box')
    const inputCatcher = await findByTestId('input-catcher')

    await user.click(typingBox)

    expect(inputCatcher).toHaveFocus()
  })

  it('should hide reminder-to-focus on click', async () => {
    const { findByTestId, queryByTestId } = render(<TypingBox />)
    const typingBox = await screen.findByTestId('typing-box')
    const user = userEvent.setup()

    const reminderToFocusBeforeClick = queryByTestId('reminder-to-focus')
    expect(reminderToFocusBeforeClick).not.toBeInTheDocument()

    const inputCatcher = await findByTestId('input-catcher')
    vi.useFakeTimers()
    inputCatcher.blur()
    vi.advanceTimersToNextTimer()

    vi.useRealTimers()

    user.click(typingBox)

    const reminderToFocusAfterClick = queryByTestId('reminder-to-focus')
    expect(reminderToFocusAfterClick).not.toBeInTheDocument()
  })
})
