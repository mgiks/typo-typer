import { useEffect, useRef, useState } from 'react'
import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'
import TextContainer from './text-container/TextContainer'
import FocusReminder from './focus-reminder/FocusReminder'
import { useAppDispatch, useAppSelector } from '../../hooks'
import {
  increaseCorrectKeysPressed,
  increaseTotalKeysPressed,
} from '../../slices/typingData.slice'
import {
  playerFinishedTyping,
  playerStartedTyping,
  playerStatusInitialState,
} from '../../slices/playerStatus.slice'
import {
  fetchText,
  setIncorrectTextStartIndexTo,
  setLastTypedIndexTo,
  textDataInitialState,
} from '../../slices/textData.slice'

export type TypingBoxProps = {
  detachStateStore?: boolean
  forcedText?: string
}

function TypingBox(
  { detachStateStore, forcedText: initialText }: TypingBoxProps,
) {
  const [isFocused, setIsFocused] = useState(true)
  const [showFocusReminder, setShowFocusReminder] = useState(false)

  const prevLastTypedIndex = useRef(-1)
  const typingBoxRef = useRef<HTMLDivElement>(null)
  const inputCatcherRef = useRef<HTMLTextAreaElement>(null)

  const text = initialText ??
    (detachStateStore
      ? textDataInitialState.text
      : useAppSelector((state) => state.textData.text))
  const lastTypedIndex = detachStateStore
    ? textDataInitialState.lastTypedIndex
    : useAppSelector((state) => state.textData.lastTypedIndex)
  const incorrectTextStartIndex = detachStateStore
    ? textDataInitialState.incorrectTextStartIndex
    : useAppSelector((state) => state.textData.incorrectTextStartIndex)
  const hasPlayerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)
  const hasPlayerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)
  const dispatch = detachStateStore ? () => {} : useAppDispatch()

  const setLastTypedIndex = (i: number) => {
    dispatch(setLastTypedIndexTo(i))
  }

  const setIncorrectTextStartIndex = (i: number) => {
    dispatch(setIncorrectTextStartIndexTo(i))
  }

  useEffect(() => {
    if (!initialText) dispatch(fetchText())

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
