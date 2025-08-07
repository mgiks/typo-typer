import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import ReminderToFocus from './reminder-to-focus/ReminderToFocus'

const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = {
  text: string
}

function TypingBox() {
  const [text, setText] = useState('')
  const [lastTypedIndex, setLastTypedIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [isFocused, setIsFocused] = useState(true)
  const [showReminderToFocus, setShowReminderToFocus] = useState(false)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))
  }, [])

  return (
    <div
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
        setShowReminderToFocus={setShowReminderToFocus}
      />
      <ReminderToFocus
        visible={showReminderToFocus}
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
