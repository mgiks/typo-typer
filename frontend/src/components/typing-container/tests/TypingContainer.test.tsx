import '@testing-library/jest-dom'
import { fireEvent, render, screen } from '@testing-library/react'
import TypingContainer from '../TypingContainer'
import * as typingStatsModule from '../../../stores/TypingStatsStore.tsx'
import { useState } from 'react'

describe('TypingContainer', () => {
  it('should display child components', () => {
    render(<TypingContainer />)

    expect(screen.getByTestId(/typing-area/)).toBeInTheDocument()
    expect(screen.getByTestId(/text-area/)).toBeInTheDocument()
  })

  it("should focus 'typing-area' on keypress", () => {
    render(<TypingContainer />)

    const typingArea = screen.getByTestId(/typing-area/)
    typingArea.blur()
    expect(typingArea).not.toHaveFocus()

    fireEvent.keyPress(document)

    expect(typingArea).toHaveFocus()
  })

  it("should focus 'typing-area' when 'isDoneTyping' is true", () => {
    render(<TestWrapper />)

    const typingArea = screen.getByTestId(/typing-area/)
    typingArea.blur()
    expect(typingArea).not.toHaveFocus()

    fireEvent.click(screen.getByText('Trigger'))

    expect(typingArea).toHaveFocus()
  })
})

function TestWrapper() {
  const [isDoneTyping, setIsDoneTyping] = useState(false)

  vi.spyOn(typingStatsModule, 'useIsDoneTyping').mockImplementation(() =>
    isDoneTyping
  )

  return (
    <>
      <button onClick={() => setIsDoneTyping(true)}>Trigger</button>
      <TypingContainer />
    </>
  )
}

function mockedTypingArea(
  { ref }: { ref: React.RefObject<HTMLTextAreaElement | null> },
) {
  return <textarea data-testid='typing-area' ref={ref} />
}

function mockedTextArea(
  { typingContainerRef }: {
    typingContainerRef: React.RefObject<HTMLDivElement | null>
  },
) {
  // Needed to suppress 'unused' warnings
  typingContainerRef

  return <div data-testid='text-area' tabIndex={-1} />
}

vi.mock('../TypingArea', () => {
  return {
    default: mockedTypingArea,
  }
})

vi.mock('../TextArea', () => {
  return {
    default: mockedTextArea,
  }
})

vi.mock('../../../stores/TypingStatsStore', () => ({
  useIsDoneTyping: () => true,
}))
