import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import TextContainer from './TextContainer.tsx'
import { useState } from 'react'

describe('TextContainer', async () => {
  it('should be in the document', async () => {
    render(
      <TextContainer
        lastTypedLetterIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should display passed text', async () => {
    const text = 'Test text.'

    render(
      <TextContainer
        lastTypedLetterIndex={-1}
        incorrectTextStartIndex={-1}
        text={text}
      />,
    )

    expect(await screen.findByTestId('text-container')).toHaveTextContent(text)
  })

  it('should include cursor', async () => {
    const { findByTestId } = render(
      <TextContainer
        lastTypedLetterIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('cursor')).toBeInTheDocument()
  })

  it('should include correct text section', async () => {
    const { findByTestId } = render(
      <TextContainer
        lastTypedLetterIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('correct-text')).toBeInTheDocument()
  })

  it('should include incorrect text section', async () => {
    const { findByTestId } = render(
      <TextContainer
        lastTypedLetterIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('incorrect-text')).toBeInTheDocument()
  })

  it('should put incorrect and correct text inside corresponding sections', async () => {
    const text = 'Test text.'

    const { findByTestId } = render(
      <TextContainer
        text={text}
        lastTypedLetterIndex={text.length - 1}
        incorrectTextStartIndex={4}
      />,
    )

    expect(await findByTestId('correct-text')).toHaveTextContent(/^Test$/)
    expect(await findByTestId('incorrect-text')).toHaveTextContent(
      /^ text\.$/,
      { normalizeWhitespace: false },
    )
  })

  it('should focus ref passed as a prop when clicking on the component', async () => {
    const mockedRefToFocus = { current: { focus: () => null } }
    const spyMockedRefToFocus = vi.spyOn(mockedRefToFocus.current, 'focus')

    function Wrapper() {
      const [_, setFocused] = useState(false)
      return (
        <TextContainer
          text=''
          lastTypedLetterIndex={-1}
          incorrectTextStartIndex={-1}
          focusSetter={setFocused}
          refToFocus={mockedRefToFocus}
        />
      )
    }

    render(<Wrapper />)

    const textContainer = await screen.findByTestId('text-container')

    expect(textContainer).toBeInTheDocument()

    const user = userEvent.setup()
    await user.click(textContainer)

    expect(spyMockedRefToFocus).toHaveBeenCalledOnce()
  })
})
