import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'

type TextResponse = {
  text: string
}

function TypingBox() {
  const [text, setText] = useState('')
  const [correctText, setCorrectText] = useState('')
  const [incorrectText, setIncorrectText] = useState('')
  const [remainingText, setRemainingText] = useState('')
  const inputRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    fetch(import.meta.env.VITE_BACKEND_URL + '/api/texts')
      .then((response) => response.json())
      .then((data: TextResponse) => {
        setText(data.text)
        setRemainingText(data.text)
      })
  }, [])

  const onType = (typed: string) => {
    const {
      correctText,
      incorrectText,
      remainingText,
    } = evaluateTyping(text, typed)

    setCorrectText(correctText)
    setIncorrectText(incorrectText)
    setRemainingText(remainingText)
  }

  return (
    <div className='typing-box' onClick={() => inputRef.current?.focus()}>
      <TypingInput ref={inputRef} onType={onType} />
      <TextDisplay
        correctText={correctText}
        incorrectText={incorrectText}
        remainingText={remainingText}
      />
    </div>
  )
}

function TextDisplay(
  { correctText, incorrectText, remainingText }: {
    correctText: string
    incorrectText: string
    remainingText: string
  },
) {
  const cursorRef = useRef<HTMLSpanElement>(null)

  useEffect(() => {
    cursorRef.current?.scrollIntoView({ block: 'center', behavior: 'smooth' })
  })

  return (
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
  )
}

function TypingInput({ ref, onType }: {
  ref: React.RefObject<HTMLTextAreaElement | null>
  onType: (typed: string) => void
}) {
  return (
    <textarea
      autoFocus
      ref={ref}
      onChange={(e) => onType(e.target.value)}
      className='typing-box__input'
    />
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

export default TypingBox
