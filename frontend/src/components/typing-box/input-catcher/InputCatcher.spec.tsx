import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InputCatcher, { BLUR_TIMEOUT } from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be in the document', async () => {
    render(
      <InputCatcher
        ref={null}
        text={''}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
        focusSetter={() => null}
      />,
    )

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should be initially focused', async () => {
    render(
      <InputCatcher
        ref={null}
        text={''}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
        focusSetter={() => null}
      />,
    )

    expect(await screen.findByTestId('input-catcher')).toHaveFocus()
  })

  it("should set incorrect letter's index when current letter is wrong", async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        ref={null}
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
        focusSetter={() => null}
      />,
    )
    await user.keyboard('tezt')

    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)
  })

  it('should not reset incorrect letter index if set incorrect letter index is smaller', async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        ref={null}
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
        focusSetter={() => null}
      />,
    )
    await user.keyboard('tez')

    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)

    await user.keyboard('d')

    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)
  })

  it('should unset incorrect letter index when last typed letter index is behind it', async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        ref={null}
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
        focusSetter={() => null}
      />,
    )
    await user.keyboard('tezt')

    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)

    await user.keyboard('{Backspace>2/}')

    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(-1)
  })

  it('should update last typed letter index', async () => {
    const spyLastTypedLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        ref={null}
        text={'test'}
        lastTypedLetterIndexSetter={spyLastTypedLetterIndexSetter}
        incorrectTextStartIndexSetter={() => null}
        focusSetter={() => null}
      />,
    )

    await user.keyboard('tes')

    expect(spyLastTypedLetterIndexSetter).toHaveReturnedWith(2)
  })

  it("should set focused to false after 'BLUR_TIMEOUT' when blurred", async () => {
    const spyFocusedSetter = vi.fn((i: boolean) => i)
    render(
      <InputCatcher
        ref={null}
        text=''
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
        focusSetter={spyFocusedSetter}
      />,
    )
    const inputCatcher = await screen.findByTestId('input-catcher')

    vi.useFakeTimers()
    inputCatcher.blur()
    vi.advanceTimersByTime(BLUR_TIMEOUT)

    vi.useRealTimers()

    expect(spyFocusedSetter).toHaveReturnedWith(false)
  })

  it('should set focused to true when focused', async () => {
    const spyFocusedSetter = vi.fn((i: boolean) => i)

    render(
      <InputCatcher
        ref={null}
        text=''
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
        focusSetter={spyFocusedSetter}
      />,
    )

    const inputCatcher = await screen.findByTestId('input-catcher')

    inputCatcher.blur()
    inputCatcher.focus()

    expect(spyFocusedSetter).toHaveReturnedWith(true)
  })
})
