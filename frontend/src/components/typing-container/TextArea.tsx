import { useEffect, useRef, useState } from 'react'
import './TextArea.css'
import { extractCorrectText } from './utils/extractCorrectText'
import { extractWrongText } from './utils/extractWrongText'
import InactivityCurtain from './InactivityCurtain'
import {
  useTextActions,
  useTextAfterCursor,
  useTextBeforeCursor,
  useTextRefreshCount,
  useWrongTextStartIndex,
} from '../../stores/TextStore'

function TextArea({
  typingContainerRef,
}: { typingContainerRef: React.RefObject<HTMLDivElement | null> }) {
  const textRefreshCount = useTextRefreshCount()
  const textBeforeCursor = useTextBeforeCursor()
  const textAfterCursor = useTextAfterCursor()
  const wrongTextStartIndex = useWrongTextStartIndex()

  const { setCorrectText } = useTextActions()

  const textAreaRef = useRef<HTMLDivElement>(null)
  const cursorRef = useRef<HTMLSpanElement>(null)
  useEffect(() => keepCursorInView(cursorRef.current, textAreaRef.current), [
    textBeforeCursor,
    textAfterCursor,
  ])

  const [isCurrentlyTyping, setIsCurrentlyTyping] = useState(false)
  const timeoutRef = useRef<NodeJS.Timeout>(null)
  useEffect(() => {
    if (timeoutRef.current) clearTimeout(timeoutRef.current)

    setIsCurrentlyTyping(true)

    timeoutRef.current = setTimeout(() => {
      setIsCurrentlyTyping(false)
    }, 750)
  }, [textBeforeCursor])

  const correctText = extractCorrectText(textBeforeCursor, wrongTextStartIndex)
  useEffect(() => {
    setCorrectText(correctText)
  })

  useEffect(() => {
    if (!textAreaRef.current) return

    const textArea = textAreaRef.current

    textArea.style.animation = 'none'

    // Causes a reflow to reset the animation
    textArea.offsetHeight

    textArea.style.animation = ''
  }, [textRefreshCount])

  useEffect(() => {
    if (typingContainerRef.current && textAreaRef.current) {
      const textArea = textAreaRef.current
      const typingContainer = typingContainerRef.current

      // It's easier to work with integer line height
      const lineHeight = Math.round(
        parseFloat(getComputedStyle(textArea).lineHeight),
      )
      textArea.style.lineHeight = lineHeight.toString() + 'px'

      const topAndBottomMargin = 30
      const numberOfLines = 5
      const typingContainerHeight =
        (lineHeight * numberOfLines + topAndBottomMargin).toString() + 'px'

      typingContainer.style.height = typingContainerHeight
    }
  }, [typingContainerRef.current, textAreaRef.current])

  const wrongText = extractWrongText(textBeforeCursor, wrongTextStartIndex)
  return (
    <>
      <InactivityCurtain />
      <div id='text-area' ref={textAreaRef}>
        <span id='correct-text' className='unselectable-text'>
          {correctText}
        </span>
        <span id='wrong-text' className='unselectable-text'>{wrongText}</span>
        <span
          id='cursor'
          ref={cursorRef}
          className={isCurrentlyTyping ? 'paused' : ''}
        >
        </span>
        <span id='trailing-text' className='unselectable-text'>
          {textAfterCursor}
        </span>
      </div>
    </>
  )
}

function keepCursorInView(
  cursor: HTMLSpanElement | null,
  textArea: HTMLDivElement | null,
) {
  if (!cursor || !textArea) {
    return
  }
  const cursorYPosition = cursor.getBoundingClientRect().y
  const textAreaTopYCoor = textArea.getBoundingClientRect().y
  const textAreaHeight = textArea.getBoundingClientRect().height
  const textAreaBottomYCoor = textAreaTopYCoor + textAreaHeight
  const isCursorInsideTextArea = textAreaTopYCoor < cursorYPosition &&
    cursorYPosition < textAreaBottomYCoor

  if (!isCursorInsideTextArea) {
    cursor.scrollIntoView({ behavior: 'smooth' })
  }
}

export default TextArea
