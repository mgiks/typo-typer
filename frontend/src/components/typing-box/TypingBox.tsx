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
    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))
  }, [])

  return (
    <div className='typing-box' data-testid='typing-box'>
      <InputCatcher />
      {text}
    </div>
  )
}

export default TypingBox
