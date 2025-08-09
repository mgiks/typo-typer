import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import FocusReminder from './focus-reminder/FocusReminder'

const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = {
  text: string
}

function TypingBox() {
  const [text, setText] = useState('')
  const [lastTypedIndex, setLastTypedIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [isFocused, setIsFocused] = useState(true)
  const [showFocusReminder, setShowFocusReminder] = useState(false)
  const ref = useRef<HTMLDivElement>(null)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    document.addEventListener('keyup', () => ref.current?.click())

    return document.removeEventListener('keyup', () => ref.current?.click())
  }, [])

  useEffect(() => {
    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))
  }, [])

  return (
    <div
      ref={ref}
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
        text={text}
        lastTypedIndex={lastTypedIndex}
        incorrectTextStartIndex={incorrectTextStartIndex}
        showCursor={isFocused}
      />
    </div>
  )
}

export default TypingBox
