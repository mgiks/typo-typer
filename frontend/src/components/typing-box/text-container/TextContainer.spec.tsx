import { render, screen } from '@testing-library/react'
import TextContainer, { type TextContainerProps } from './TextContainer.tsx'

describe('TextContainer', async () => {
  it('should be in the document', async () => {
    renderTextContainer()

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it("should display text from 'text' prop", async () => {
    const text = 'Test text.'

    renderTextContainer({ text: text })

    expect(await screen.findByTestId('text-container')).toHaveTextContent(text)
  })

  it('should include a section for correct text', async () => {
    const { findByTestId } = renderTextContainer()

    expect(await findByTestId('correct-text')).toBeInTheDocument()
  })

  it('should include a section for incorrect text', async () => {
    const { findByTestId } = renderTextContainer()

    expect(await findByTestId('incorrect-text')).toBeInTheDocument()
  })

  it("should include cursor when 'showCursor' prop is true", async () => {
    const { findByTestId } = renderTextContainer({ showCursor: true })

    expect(await findByTestId('cursor')).toBeInTheDocument()
  })

  it("should not include cursor when 'showCursor' prop is false", async () => {
    const { queryByTestId } = renderTextContainer()

    expect(queryByTestId('cursor')).not.toBeInTheDocument()
  })

  it('should put incorrect and correct text inside corresponding sections', async () => {
    const text = 'Test text.'

    const { findByTestId } = renderTextContainer({
      text: text,
      lastTypedIndex: text.length - 1,
      incorrectTextStartIndex: 4,
    })

    expect(await findByTestId('correct-text')).toHaveTextContent(/^Test$/)
    expect(await findByTestId('incorrect-text')).toHaveTextContent(
      /^ text\.$/,
      { normalizeWhitespace: false },
    )
  })

  it("should set cursor y position on 'lastTypedIndex' prop change", async () => {
    const setCursorYPosition = vi.fn((_: number) => {})
    const prereqProps: TextContainerProps = {
      ...defaultProps,
      showCursor: true,
      setCursorYPosition: setCursorYPosition,
    }

    const { rerender } = render(
      <TextContainer {...prereqProps} lastTypedIndex={0} />,
    )

    rerender(<TextContainer {...prereqProps} lastTypedIndex={1} />)

    expect(setCursorYPosition).toHaveBeenCalled()
  })
})

const defaultProps: TextContainerProps = {
  showCursor: false,
  lastTypedIndex: -1,
  incorrectTextStartIndex: -1,
  setCursorYPosition: (_: number) => {},
  text: '',
}

function renderTextContainer(overrides: Partial<TextContainerProps> = {}) {
  return render(<TextContainer {...defaultProps} {...overrides} />)
}
