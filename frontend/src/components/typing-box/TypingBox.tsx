import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import FocusReminder from './focus-reminder/FocusReminder'

const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = { text: string }

function TypingBox() {
  const [text, setText] = useState('')
  const [lastTypedIndex, setLastTypedIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [isFocused, setIsFocused] = useState(true)
  const [showFocusReminder, setShowFocusReminder] = useState(false)
  const [cursorYPosition, setCursorYPosition] = useState(-1)

  const scrollDistance = useRef(51)
  const prevCursorYPosition = useRef(-1)
  const typingBoxRef = useRef<HTMLDivElement>(null)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    document.addEventListener('keyup', () => typingBoxRef.current?.click())

    return document.removeEventListener(
      'keyup',
      () => typingBoxRef.current?.click(),
    )
  }, [])

  useEffect(() => {
    typingBoxRef.current?.scrollTo(0, 0)

    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))
  }, [])

  useEffect(() => {
    if (
      prevCursorYPosition.current !== -1 &&
      (cursorYPosition !== prevCursorYPosition.current)
    ) {
      // 'scrollBy' works weird for some reason, so 'scrollTo' instead
      typingBoxRef.current?.scrollTo(0, scrollDistance.current)

      scrollDistance.current += 51
    }

    prevCursorYPosition.current = cursorYPosition
  }, [cursorYPosition])

  return (
    <div
      ref={typingBoxRef}
      className='typing-box'
      data-testid='typing-box'
      onClick={() => inputCatcherRef.current?.focus()}
    >
      <InputCatcher
        ref={inputCatcherRef}
        text={text}
        setIncorrectTextStartIndex={setIncorrectTextStartIndex}
        setIsFocused={setIsFocused}
        setLastTypedIndex={setLastTypedIndex}
        setShowFocusReminder={setShowFocusReminder}
      />
      <FocusReminder
        visible={showFocusReminder}
      />
      <TextContainer
        setCursorYPosition={setCursorYPosition}
        text={text}
        lastTypedIndex={lastTypedIndex}
        incorrectTextStartIndex={incorrectTextStartIndex}
        showCursor={isFocused}
      />
    </div>
  )
}

export default TypingBox
