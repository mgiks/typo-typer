import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import FocusReminder from './focus-reminder/FocusReminder'
import { useAppDispatch, useAppSelector } from '../../hooks'
import {
  increaseCorrectKeysPressed,
  increaseTotalKeysPressed,
} from '../../slices/typingStats.slice'
import {
  playerFinishedTyping,
  playerStartedTyping,
  playerStatusInitialState,
} from '../../slices/playerStatus.slice'

export const TEXTS_URL = 'http://localhost:8000/texts'

export type GETTextResponse = { text: string }

export type TypingBoxProps = {
  detachStateStore?: boolean
  initialText?: string
}

function TypingBox({ detachStateStore, initialText }: TypingBoxProps) {
  const [text, setText] = useState(initialText ?? '')
  const [lastTypedIndex, setLastTypedIndex] = useState(-1)
  const [incorrectTextStartIndex, setIncorrectTextStartIndex] = useState(-1)
  const [isFocused, setIsFocused] = useState(true)
  const [showFocusReminder, setShowFocusReminder] = useState(false)

  const prevLastTypedIndex = useRef(-1)
  const typingBoxRef = useRef<HTMLDivElement>(null)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  const hasPlayerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)
  const hasPlayerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)
  const dispatch = detachStateStore ? () => {} : useAppDispatch()

  useEffect(() => {
    if (!initialText) {
      fetch(TEXTS_URL)
        .then((resp) => resp.json())
        .then((resp) => resp as GETTextResponse)
        .then((json) => setText(json.text))
        .catch((_) => console.error('Network error'))
    }

    const handleKeyPress = () => typingBoxRef.current?.click()

    document.addEventListener('keyup', handleKeyPress)

    return () => document.removeEventListener('keyup', handleKeyPress)
  }, [])

  useEffect(() => {
    if (text && lastTypedIndex === text.length - 1) {
      dispatch(playerFinishedTyping())
    }

    if (!hasPlayerStartedTyping && lastTypedIndex > -1) {
      dispatch(playerStartedTyping())
    }

    if (lastTypedIndex > prevLastTypedIndex.current) {
      if (incorrectTextStartIndex === -1) dispatch(increaseCorrectKeysPressed())
      dispatch(increaseTotalKeysPressed())
    }

    if (prevLastTypedIndex.current !== lastTypedIndex) {
      prevLastTypedIndex.current = lastTypedIndex
    }
  }, [lastTypedIndex, incorrectTextStartIndex])

  return (!hasPlayerFinishedTyping && (
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
