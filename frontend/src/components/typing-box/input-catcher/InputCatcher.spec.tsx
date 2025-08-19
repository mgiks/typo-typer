import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InputCatcher, {
  FOCUS_REMINDER_TIMEOUT_MS,
  type InputCatcherProps,
} from './InputCatcher.tsx'
import { TEXT_FIXTURE } from '../../../tests/fixtures.ts'

describe('InputCatcher', async () => {
  it('should be in the document', () => {
    renderInputCatcher()

    expect(screen.getByRole('textbox')).toBeInTheDocument()
  })

  it('should be initially focused', () => {
    renderInputCatcher()

    expect(screen.getByRole('textbox')).toHaveFocus()
  })

  it(
    "should set start index of incorrect text as the current letter's index when current letter is wrong",
    async () => {
      const user = userEvent.setup()
      const setIncorrectTextStartIndex = vi.fn((i: number) => i)

      renderInputCatcher({
        text: TEXT_FIXTURE,
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('Tezt')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)
    },
  )

  it(
    "should not reset start index of incorrect text when current incorrect letter's index is bigger",
    async () => {
      const user = userEvent.setup()
      const setIncorrectTextStartIndex = vi.fn((i: number) => i)

      renderInputCatcher({
        text: TEXT_FIXTURE,
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('Tez')

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
        text: TEXT_FIXTURE,
        setIncorrectTextStartIndex: setIncorrectTextStartIndex,
      })

      await user.keyboard('Tezt')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(2)

      await user.keyboard('{Backspace>2/}')

      expect(setIncorrectTextStartIndex).toHaveReturnedWith(-1)
    },
  )

  it('should update the index of the last typed letter', async () => {
    const user = userEvent.setup()
    const setLastTypedIndex = vi.fn((i: number) => i)

    renderInputCatcher({
      text: TEXT_FIXTURE,
      setLastTypedIndex: setLastTypedIndex,
    })

    await user.keyboard('tes')

    expect(setLastTypedIndex).toHaveReturnedWith(2)
  })

  it("should set 'isFocused' to false after a timeout when blurred", async () => {
    const setIsFocused = vi.fn((i: boolean) => i)
    renderInputCatcher({ setIsFocused: setIsFocused })
    const inputCatcher = screen.getByRole('textbox')

    vi.useFakeTimers()
    inputCatcher.blur()
    vi.advanceTimersByTime(FOCUS_REMINDER_TIMEOUT_MS).useRealTimers()

    expect(setIsFocused).toHaveReturnedWith(false)
  })

  it("should set 'isFocused' to true when focused", async () => {
    const setIsFocused = vi.fn((i: boolean) => i)

    renderInputCatcher({ setIsFocused: setIsFocused })
    const inputCatcher = screen.getByRole('textbox')

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
