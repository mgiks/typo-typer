import { useEffect, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'

const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = {
  text: string
}

function TypingBox() {
  const [text, setText] = useState('')

  useEffect(() => {
    try {
      fetch(TEXTS_URL)
        .then((resp) => resp.json())
        .then((resp) => resp as GETTextResponse)
        .then((json) => setText(json.text))
    } catch (e) {
    }
  })

  return (
    <div className='typing-box' data-testid='typing-box'>
      <InputCatcher />
      {text}
    </div>
  )
}

export default TypingBox
