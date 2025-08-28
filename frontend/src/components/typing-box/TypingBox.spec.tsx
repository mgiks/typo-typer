import { act, render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox, { type TypingBoxProps } from './TypingBox.tsx'
import { FOCUS_REMINDER_TIMEOUT_MS } from './input-catcher/InputCatcher.tsx'
import { TEXT_FIXTURE } from '../../tests/fixtures.ts'
import { TEXTS_URL } from '../../slices/textData.slice.ts'
import { renderWithProviders } from '../../tests/utils.tsx'

const FOCUS_REMINDER_TEXT = /click here or press any key to focus/i

describe('TypingBox', async () => {
  it('should be in the document', () => {
    renderTypingBox()

    expect(screen.getByRole('region')).toBeInTheDocument()
  })

  it('should display fetched text on initial render', async () => {
    const server = setupServer(
      http.get(TEXTS_URL, () => HttpResponse.json({ text: TEXT_FIXTURE })),
    )
    server.listen()

    renderWithProviders(<TypingBox />)

    expect(await screen.findByText('Test text.')).toBeInTheDocument()

    server.close()
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

const defaultProps: TypingBoxProps = {
  detachStateStore: true,
  initialText: 'Placeholder',
}

function renderTypingBox(overrides?: Partial<TypingBoxProps>) {
  return render(<TypingBox {...defaultProps} {...overrides} />)
}
