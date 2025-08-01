import { useEffect, useState } from 'react'
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
  const [lastTypedLetterIndex, setLastTypedLetterIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [focused, setFocused] = useState(true)

  useEffect(() => {
    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))
  }, [])

  return (
    <div className='typing-box' data-testid='typing-box'>
      <InputCatcher
        text={text}
        lastTypedLetterIndexSetter={setLastTypedLetterIndex}
        incorrectTextStartIndexSetter={setIncorrectTextStartIndex}
        focusedSetter={setFocused}
      />
      <ReminderToFocus focused={focused} focusSetter={setFocused} />
      <TextContainer
        text={text}
        lastTypedLetterIndex={lastTypedLetterIndex}
        incorrectTextStartIndex={incorrectTextStartIndex}
      />
    </div>
  )
}

export default TypingBox
