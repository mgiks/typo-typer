import { act, render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox, { TEXTS_URL } from './TypingBox.tsx'
import { FOCUS_REMINDER_TIMEOUT_MS } from './input-catcher/InputCatcher.tsx'
import { TEXT_FIXTURE } from '../../tests/fixtures.ts'
import { Provider } from 'react-redux'
import { store } from '../../store.ts'

const FOCUS_REMINDER_TEXT = /click here or press any key to focus/i

const handlers = [
  http.get(TEXTS_URL, () => HttpResponse.json({ text: TEXT_FIXTURE })),
]

const server = setupServer(...handlers)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('TypingBox', async () => {
  it('should be in the document', () => {
    renderTypingBox()

    expect(screen.getByRole('region')).toBeInTheDocument()
  })

  it('should display fetched text on initial render', async () => {
    renderTypingBox()

    expect(await screen.findByText('Test text.')).toBeInTheDocument()
  })

  it('should show focus reminder when input catcher is blurred', () => {
    const { getByRole } = renderTypingBox()
    const inputCatcher = getByRole('textbox')

    expect(screen.queryByText(FOCUS_REMINDER_TEXT)).not.toBeInTheDocument()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(screen.getByText(FOCUS_REMINDER_TEXT)).toBeInTheDocument()
  })

  it('should focus input catcher on click', async () => {
    const { getByRole } = renderTypingBox()
    const typingBox = screen.getByRole('region')
    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    await user.click(typingBox)

    expect(inputCatcher).toHaveFocus()
  })

  it('should focus input catcher on keypress', async () => {
    const { getByRole } = renderTypingBox()
    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(inputCatcher).not.toHaveFocus()

    await user.keyboard('t')

    expect(inputCatcher).toHaveFocus()
  })

  it('should hide focus reminder on click', async () => {
    const { getByRole } = renderTypingBox()
    const typingBox = screen.getByRole('region')
    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    expect(screen.queryByText(FOCUS_REMINDER_TEXT)).not.toBeInTheDocument()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()
    })

    expect(screen.queryByText(FOCUS_REMINDER_TEXT)).toBeInTheDocument()

    await user.click(typingBox)

    expect(screen.queryByText(FOCUS_REMINDER_TEXT)).not.toBeInTheDocument()
  })
})

function renderTypingBox() {
  return render(
    <Provider store={store}>
      <TypingBox />
    </Provider>,
  )
}
