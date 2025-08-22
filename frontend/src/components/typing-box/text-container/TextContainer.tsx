import { useEffect, useRef } from 'react'
import './TextContainer.scss'
import Cursor from './cursor/Cursor'

export type TextContainerProps = {
  showCursor: boolean
  text: string
  lastTypedIndex: number
  incorrectTextStartIndex: number
}

function TextContainer(
  { text, showCursor, lastTypedIndex, incorrectTextStartIndex }:
    TextContainerProps,
) {
  let correctText = ''
  let incorrectText = ''

  if (incorrectTextStartIndex === -1) {
    correctText = text.slice(0, lastTypedIndex + 1)
  } else {
    correctText = text.slice(0, incorrectTextStartIndex)
    incorrectText = text.slice(
      incorrectTextStartIndex,
      lastTypedIndex + 1,
    )
  }

  const textContainerRef = useRef<HTMLDivElement>(null)
  const cursorRef = useRef<HTMLSpanElement>(null)

  const scrollCursorIntoView = () => {
    cursorRef.current?.scrollIntoView({ behavior: 'instant', block: 'center' })
  }

  useEffect(() => {
    if (!textContainerRef.current) return

    const observer = new ResizeObserver(() => scrollCursorIntoView())

    observer.observe(textContainerRef.current)

    return () => observer.disconnect()
  }, [])

  useEffect(() => {
    const cursorRect = cursorRef.current?.getBoundingClientRect()
    const textContainerRect = textContainerRef.current?.getBoundingClientRect()

    if (!cursorRect || !textContainerRect) return

    const textContainerCenter = textContainerRect.top +
      textContainerRect.height / 2
    const cursorCenter = cursorRect.top + cursorRect.height / 2

    // Needed to prevent tiny scroll adjustments when this is not necessary
    const isCursorOffCentered = Math.abs(textContainerCenter - cursorCenter) > 5

    if (isCursorOffCentered) scrollCursorIntoView()
  }, [lastTypedIndex])

  return (
    <div
      ref={textContainerRef}
      className='text-container'
      role='status'
    >
      <span className='text-container__text_correct' aria-label='Correct text'>
        {correctText}
      </span>
      <span
        className='text-container__text_incorrect'
        aria-label='Incorrect text'
      >
        {incorrectText}
      </span>
      <Cursor
        ref={cursorRef}
        visible={showCursor}
      />
      <span>{text.slice(lastTypedIndex + 1)}</span>
    </div>
  )
}

export default TextContainer
