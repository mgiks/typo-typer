import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InputCatcher, {
  FOCUS_REMINDER_TIMEOUT_MS,
  type InputCatcherProps,
} from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be in the document', async () => {
    renderInputCatcher()

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should be initially focused', async () => {
    renderInputCatcher()

    expect(await screen.findByTestId('input-catcher')).toHaveFocus()
  })

  it(
    "should set start index of incorrect text as the current letter's index when current letter is wrong",
    async () => {
      const user = userEvent.setup()
      const setIncorrectTextStartIndex = vi.fn((i: number) => i)

      renderInputCatcher({
        text: 'test',
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('tezt')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)
    },
  )

  it(
    "should not reset start index of incorrect text when current incorrect letter's index is bigger",
    async () => {
      const user = userEvent.setup()
      const setIncorrectTextStartIndex = vi.fn((i: number) => i)

      renderInputCatcher({
        text: 'test',
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('tez')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)

      await user.keyboard('d')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)
    },
  )

  it(
    "should unset start index of incorrect text when last-typed letter's index is smaller",
    async () => {
      const user = userEvent.setup()
      const setIncorrectTextStartIndex = vi.fn((i: number) => i)

      renderInputCatcher({
        text: 'test',
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('tezt')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)

      await user.keyboard('{Backspace>2/}')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(-1)
    },
  )

  it('should update last typed letter index', async () => {
    const user = userEvent.setup()
    const setLastTypedIndex = vi.fn((i: number) => i)

    renderInputCatcher({
      text: 'test',
      setLastTypedIndex: setLastTypedIndex,
    })

    await user.keyboard('tes')

    expect(setLastTypedIndex).toHaveReturnedWith(2)
  })

  it("should set 'isFocused' to false after a timeout when blurred", async () => {
    const setIsFocused = vi.fn((i: boolean) => i)

    renderInputCatcher({ setIsFocused: setIsFocused })
    const inputCatcher = await screen.findByTestId('input-catcher')

    vi.useFakeTimers()
    inputCatcher.blur()
    vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()

    expect(setIsFocused).toHaveReturnedWith(false)
  })

  it("should set 'isFocused' to true when focused", async () => {
    const setIsFocused = vi.fn((i: boolean) => i)

    renderInputCatcher({ setIsFocused: setIsFocused })
    const inputCatcher = await screen.findByTestId('input-catcher')

    vi.useFakeTimers()
    inputCatcher.blur()
    vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()

    inputCatcher.focus()

    expect(setIsFocused).toHaveReturnedWith(true)
  })
})

const setterStub = () => {}

const defaultProps: InputCatcherProps = {
  ref: null,
  text: '',
  setIncorrectTextStartIndex: setterStub,
  setIsFocused: setterStub,
  setLastTypedIndex: setterStub,
  setShowFocusReminder: setterStub,
}

function renderInputCatcher(overrides: Partial<InputCatcherProps> = {}) {
  return render(<InputCatcher {...defaultProps} {...overrides} />)
}
