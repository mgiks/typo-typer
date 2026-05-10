import { useEffect, useRef, useState } from 'react'
import './TypingPage.scss'
import { API_URL } from '../App'

type Text = {
  data: {
    id: string
    content: string
  }
}

function TypingPage() {
  const [text, setText] = useState('')
  const [correctText, setCorrectText] = useState('')
  const [incorrectText, setIncorrectText] = useState('')
  const [remainingText, setRemainingText] = useState('')
  const inputRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    fetch(API_URL + '/v1/texts/random')
      .then((response) => response.json())
      .then((payload: Text) => {
        setText(payload.data.content)
        setRemainingText(payload.data.content)
      })
  }, [])

  const onType = (typed: string) => {
    const {
      correctText,
      incorrectText,
      remainingText,
    } = evaluateTyping(
      text,
      typed,
    )

    setCorrectText(correctText)
    setIncorrectText(incorrectText)
    setRemainingText(remainingText)
  }

  const cursorRef = useRef<HTMLSpanElement>(null)

  useEffect(() => {
    cursorRef.current?.scrollIntoView({ block: 'center', behavior: 'smooth' })
  })

  return (
    <div className='typing-box' onClick={() => inputRef.current?.focus()}>
      <textarea
        autoFocus
        ref={inputRef}
        onChange={(e) => onType(e.target.value)}
        className='typing-box__input'
      />
      <div className='typing-box__text'>
        <span className='typing-box__text--correct' content={correctText}>
          {correctText}
        </span>
        <span className='typing-box__text--incorrect' content={incorrectText}>
          {incorrectText}
        </span>
        <span ref={cursorRef} className='typing-box__cursor' />
        {remainingText}
      </div>
    </div>
  )
}

function evaluateTyping(text: string, typed: string) {
  let incorrectTextIndex = -1

  for (let i = 0; i < typed.length; ++i) {
    if (text.at(i) !== typed.at(i)) {
      incorrectTextIndex = i
      break
    }
  }

  const correctText = incorrectTextIndex > -1
    ? typed.slice(0, incorrectTextIndex)
    : typed

  const incorrectText = incorrectTextIndex > -1
    ? text.slice(incorrectTextIndex, typed.length)
    : ''

  const remainingText = text.slice(typed.length)

  return {
    correctText: correctText,
    incorrectText: incorrectText,
    remainingText: remainingText,
  }
}

export default TypingPage
