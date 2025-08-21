import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import FocusReminder from './focus-reminder/FocusReminder'
import { useAppDispatch, useAppSelector } from '../../hooks'
import { userStartedTyping } from '../../slices/isUserTyping.slice'
import {
  increaseCorrectKeysPressed,
  increaseTotalKeysPressed,
} from '../../slices/typingStats.slice'

export const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = { text: string }

function TypingBox() {
  const [text, setText] = useState('')
  const [lastTypedIndex, setLastTypedIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [isFocused, setIsFocused] = useState(true)
  const [showFocusReminder, setShowFocusReminder] = useState(false)
  const [userFinishedTyping, setUserFinishedTyping] = useState(false)

  const prevLastTypedIndex = useRef(-1)
  const typingBoxRef = useRef<HTMLDivElement>(null)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  const isUserTyping = useAppSelector((state) => state.isUserTyping.value)
  const dispatch = useAppDispatch()

  useEffect(() => {
    fetch(TEXTS_URL)
      .then((resp) => resp.json())
      .then((resp) => resp as GETTextResponse)
      .then((json) => setText(json.text))
      .catch((_) => console.error('Network error'))

    const handleKeyPress = () => typingBoxRef.current?.click()

    document.addEventListener('keyup', handleKeyPress)

    return () => document.removeEventListener('keyup', handleKeyPress)
  }, [])

  useEffect(() => {
    if (text && lastTypedIndex === text.length - 1) setUserFinishedTyping(true)

    if (!isUserTyping && lastTypedIndex > -1) dispatch(userStartedTyping())

    if (lastTypedIndex > prevLastTypedIndex.current) {
      if (incorrectTextStartIndex === -1) dispatch(increaseCorrectKeysPressed())
      dispatch(increaseTotalKeysPressed())
    }

    if (prevLastTypedIndex.current !== lastTypedIndex) {
      prevLastTypedIndex.current = lastTypedIndex
    }
  }, [lastTypedIndex, incorrectTextStartIndex])

  return (!userFinishedTyping && (
    <div
      ref={typingBoxRef}
      className='typing-box'
      role='region'
      aria-label='Typing Box'
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
  ))
}

export default TypingBox
