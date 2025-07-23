import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import InputCatcher from './InputCatcher.tsx'

describe('InputCatcher', async () => {
  it('should be in the document', async () => {
    render(
      <InputCatcher
        text={''}
        incorrectLetterIndexSetter={() => null}
      />,
    )

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })

  it('should be initially focused', async () => {
    render(
      <InputCatcher
        text={''}
        incorrectLetterIndexSetter={() => null}
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
        incorrectLetterIndexSetter={spyIncorrectLetterIndexSetter}
      />,
    )
    await user.keyboard('tezt')

    expect(spyIncorrectLetterIndexSetter).toHaveBeenCalledOnce()
    expect(spyIncorrectLetterIndexSetter).toHaveReturnedWith(2)
  })
})
