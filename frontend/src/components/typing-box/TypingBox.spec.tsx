import { act, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import userEvent from '@testing-library/user-event'
import TypingBox from './TypingBox.tsx'
import { FOCUS_REMINDER } from './focus-reminder/FocusReminder'
import { TEXT_FIXTURE } from '../../tests/fixtures.ts'
import { TEXTS_URL } from '../../slices/textData.slice.ts'
import { renderWithProviders } from '../../tests/utils.tsx'

describe('TypingBox', async () => {
  const server = setupServer(
    http.get(TEXTS_URL, () => HttpResponse.json({ text: TEXT_FIXTURE })),
  )

  beforeAll(() => server.listen())
  afterEach(() => server.resetHandlers())
  afterAll(() => server.close())

  it(
    'should display fetched text on initial render',
    async () => {
      renderWithProviders(<TypingBox />)

      expect(await screen.findByText('Test text.')).toBeInTheDocument()
    },
  )

  it(
    'should show focus reminder when input catcher is blurred',
    async () => {
      const { getByRole } = renderWithProviders(<TypingBox />)

      const inputCatcher = getByRole('textbox')

      expect(screen.queryByText(FOCUS_REMINDER.TEXT)).not.toBeInTheDocument()

      act(() => {
        vi.useFakeTimers()
        inputCatcher.blur()
        vi.runAllTimersAsync()
      })

      expect(await screen.findByText(FOCUS_REMINDER.TEXT)).toBeInTheDocument()
    },
  )

  it('should focus input catcher on click', async () => {
    const { getByRole } = renderWithProviders(<TypingBox />)

    const typingBox = screen.getByRole('region')
    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.runAllTimersAsync()
    })

    await user.click(typingBox)

    expect(inputCatcher).toHaveFocus()
  })

  it('should focus input catcher on keypress', async () => {
    const { getByRole } = renderWithProviders(<TypingBox />)

    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.advanceTimersByTime(FOCUS_REMINDER.TIMEOUT_MS).useRealTimers()
    })

    expect(inputCatcher).not.toHaveFocus()

    await user.keyboard('t')

    expect(inputCatcher).toHaveFocus()
  })

  it('should hide focus reminder on click', async () => {
    const { getByRole } = renderWithProviders(<TypingBox />)

    const typingBox = screen.getByRole('region')
    const inputCatcher = getByRole('textbox')
    const user = userEvent.setup()

    expect(screen.queryByText(FOCUS_REMINDER.TEXT)).not.toBeInTheDocument()

    act(() => {
      vi.useFakeTimers()
      inputCatcher.blur()
      vi.runAllTimersAsync()
    })

    expect(await screen.findByText(FOCUS_REMINDER.TEXT)).toBeInTheDocument()

    await user.click(typingBox)

    expect(screen.queryByText(FOCUS_REMINDER.TEXT)).not.toBeInTheDocument()
  })
})
