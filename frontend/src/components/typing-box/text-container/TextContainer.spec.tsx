import { render, screen } from '@testing-library/react'
import TextContainer from './TextContainer.tsx'

describe('TextContainer', async () => {
  it('should be in the document', async () => {
    render(
      <TextContainer
        showCursor={false}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it("should display text from 'text' prop", async () => {
    const text = 'Test text.'

    render(
      <TextContainer
        showCursor={false}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text={text}
      />,
    )

    expect(await screen.findByTestId('text-container')).toHaveTextContent(text)
  })

  it('should include a section for correct text', async () => {
    const { findByTestId } = render(
      <TextContainer
        showCursor={false}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('correct-text')).toBeInTheDocument()
  })

  it('should include a section for incorrect text', async () => {
    const { findByTestId } = render(
      <TextContainer
        showCursor={false}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('incorrect-text')).toBeInTheDocument()
  })

  it("should include cursor when 'showCursor' prop is true", async () => {
    const { findByTestId } = render(
      <TextContainer
        showCursor={true}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(await findByTestId('cursor')).toBeInTheDocument()
  })

  it("should not include cursor when 'showCursor' prop is false", async () => {
    const { queryByTestId } = render(
      <TextContainer
        showCursor={false}
        lastTypedIndex={-1}
        incorrectTextStartIndex={-1}
        text=''
      />,
    )

    expect(queryByTestId('cursor')).not.toBeInTheDocument()
  })

  it('should put incorrect and correct text inside corresponding sections', async () => {
    const text = 'Test text.'

    const { findByTestId } = render(
      <TextContainer
        showCursor={false}
        text={text}
        lastTypedIndex={text.length - 1}
        incorrectTextStartIndex={4}
      />,
    )

    expect(await findByTestId('correct-text')).toHaveTextContent(/^Test$/)
    expect(await findByTestId('incorrect-text')).toHaveTextContent(
      /^ text\.$/,
      { normalizeWhitespace: false },
    )
  })
})
