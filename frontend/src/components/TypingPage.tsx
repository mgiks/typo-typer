import { useEffect, useRef, useState } from 'react'
import './TypingPage.scss'
import { BACKEND_URL } from '../App'

type TextMsg = {
  data: {
    text: string
  }
}

function TypingPage() {
  const [text, setText] = useState('')

  let wsConn: WebSocket
  const connectWS = () => {
    wsConn = new WebSocket(BACKEND_URL + '/v1/matchmaking/pool')
    wsConn.addEventListener(
      'open',
      () => wsConn.send(JSON.stringify({ username: 'jfd' })),
    )
    wsConn.addEventListener(
      'message',
      (rawMsg) => {
        console.log('Got message')
        const msg = JSON.parse(rawMsg.data) as TextMsg
        console.log(msg)
        setText(msg.data.text)
        console.log(msg.data.text)
      },
    )
  }

  return (
    <>
      <TypingBox text={text} />
      <button onClick={connectWS}>Play</button>
    </>
  )
}

function TypingBox({ text }: { text: string }) {
  const [correctText, setCorrectText] = useState('')
  const [incorrectText, setIncorrectText] = useState('')
  const [remainingText, setRemainingText] = useState('')
  const inputRef = useRef<HTMLTextAreaElement>(null)

  useEffect(() => {
    setRemainingText(text)
  }, [text])

  const evaluateTyping = (typed: string) => {
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

    setCorrectText(correctText)
    setIncorrectText(incorrectText)
    setRemainingText(remainingText)
  }

  const cursorRef = useRef<HTMLSpanElement>(null)

  useEffect(() => {
    cursorRef.current?.scrollIntoView({ block: 'center', behavior: 'smooth' })
  })

  return (
    <section
      className='typing-box'
      onClick={() => inputRef.current?.focus()}
    >
      <textarea
        autoFocus
        ref={inputRef}
        onChange={(e) => evaluateTyping(e.target.value)}
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
    </section>
  )
}

export default TypingPage
