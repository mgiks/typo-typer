import { render, screen } from '@testing-library/react'
import TextContainer, { type TextContainerProps } from './TextContainer.tsx'
import { TEXT_FIXTURE } from '../../../tests/fixtures.ts'

describe('TextContainer', async () => {
  it('should be in the document', () => {
    renderTextContainer()

    expect(screen.getByRole('status')).toBeInTheDocument()
  })

  it("should display text from 'text' prop", () => {
    renderTextContainer({ text: TEXT_FIXTURE })

    expect(screen.getByRole('status')).toHaveTextContent(TEXT_FIXTURE)
  })

  it('should include a section for correct text', () => {
    const { getByLabelText } = renderTextContainer()

    expect(getByLabelText(/^correct text$/i)).toBeInTheDocument()
  })

  it('should include a section for incorrect text', () => {
    const { getByLabelText } = renderTextContainer()

    expect(getByLabelText(/^incorrect text$/i)).toBeInTheDocument()
  })

  it("should include cursor when 'showCursor' prop is true", () => {
    const { getByLabelText } = renderTextContainer({ showCursor: true })

    expect(getByLabelText(/cursor/i)).toBeInTheDocument()
  })

  it("should not include cursor when 'showCursor' prop is false", () => {
    const { queryByLabelText } = renderTextContainer()

    expect(queryByLabelText(/cursor/i)).not.toBeInTheDocument()
  })

  it('should put incorrect and correct text inside corresponding sections', () => {
    const { getByLabelText } = renderTextContainer({
      text: TEXT_FIXTURE,
      lastTypedIndex: TEXT_FIXTURE.length - 1,
      incorrectTextStartIndex: 4,
    })

    expect(getByLabelText(/^correct text$/i)).toHaveTextContent(/^Test$/)
    expect(getByLabelText(/^incorrect text$/i)).toHaveTextContent(
      /^ text\.$/,
      { normalizeWhitespace: false },
    )
  })
})

const defaultProps: TextContainerProps = {
  showCursor: false,
  lastTypedIndex: -1,
  incorrectTextStartIndex: -1,
  text: '',
  cursorRef: null,
}

function renderTextContainer(overrides: Partial<TextContainerProps> = {}) {
  return render(<TextContainer {...defaultProps} {...overrides} />)
}
