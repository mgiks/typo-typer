import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InputCatcher from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be in the document', async () => {
    render(
      <InputCatcher
        text={''}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
      />,
    )

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should be initially focused', async () => {
    render(
      <InputCatcher
        text={''}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={() => null}
      />,
    )

    expect(await screen.findByTestId('input-catcher')).toHaveFocus()
  })

  it('should set incorrect letter index when current letter is wrong', async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
      />,
    )
    await user.keyboard('tezt')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledOnce()
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)
  })

  it('should not reset incorrect letter index if set incorrect letter index is smaller', async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
      />,
    )
    await user.keyboard('tez')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledOnce()
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)

    await user.keyboard('d')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledOnce()
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)
  })

  it('should unset incorrect letter index when last typed letter index is behind it', async () => {
    const spyIncorrectLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        text={'test'}
        lastTypedLetterIndexSetter={() => null}
        incorrectTextStartIndexSetter={spyIncorrectLetterIndexSetter}
      />,
    )
    await user.keyboard('tezt')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledOnce()
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)

    await user.keyboard('{Backspace>2/}')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledTimes(2)
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(-1)
  })

  it('should update last typed letter index', async () => {
    const spyLastTypedLetterIndexSetter = vi.fn((i: number) => i)
    const user = userEvent.setup()

    render(
      <InputCatcher
        text={'test'}
        lastTypedLetterIndexSetter={spyLastTypedLetterIndexSetter}
        incorrectTextStartIndexSetter={() => null}
      />,
    )

    await user.keyboard('tes')

    expect(spyLastTypedLetterIndexSetter).toHaveBeenCalledTimes(3)
    expect(spyLastTypedLetterIndexSetter).toHaveReturnedWith(2)
  })
})
