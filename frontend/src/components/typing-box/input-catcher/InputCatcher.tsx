import React, { useEffect, useRef, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../../../hooks'
import { userStartedTyping } from '../../../slices/isUserTyping.slice'

export const FOCUS_REMINDER_TIMEOUT_MS = 750

type Setter<T> = (i: T) => void

export type InputCatcherProps = {
  text: string
  setIsFocused: Setter<boolean>
  setLastTypedIndex: Setter<number>
  setShowFocusReminder: Setter<boolean>
  setIncorrectTextStartIndex: Setter<number>
  ref: React.Ref<HTMLTextAreaElement | null>
}

function InputCatcher(
  {
    ref,
    text,
    setIsFocused,
    setLastTypedIndex,
    setShowFocusReminder,
    setIncorrectTextStartIndex,
  }: InputCatcherProps,
) {
  const [input, setInput] = useState('')
  const [lastIncorrectLetterIndex, setLastIncorrectLetterIndex] = useState(-1)
  const timeOutRef = useRef(-1)

  const isUserTyping = useAppSelector((state) => state.isUserTyping.value)
  const dispatch = useAppDispatch()

  useEffect(() => {
    if (!isUserTyping && input) dispatch(userStartedTyping())

    const inputLastIndex = input.length - 1
    setLastTypedIndex(inputLastIndex)

    if (inputLastIndex === -1) {
      setLastIncorrectLetterIndex(-1)
      setIncorrectTextStartIndex(-1)
      return
    }

    if (text[inputLastIndex] !== input.at(-1)) {
      if (
        lastIncorrectLetterIndex === -1 ||
        inputLastIndex < lastIncorrectLetterIndex
      ) {
        setIncorrectTextStartIndex(inputLastIndex)
        setLastIncorrectLetterIndex(inputLastIndex)
      }
    } else if (inputLastIndex < lastIncorrectLetterIndex) {
      setIncorrectTextStartIndex(-1)
      setLastIncorrectLetterIndex(-1)
    }
  }, [input])

  return (
    <textarea
      ref={ref}
      className='typing-box__input-catcher'
      onInput={(event) => setInput(event.currentTarget.value)}
      onBlur={() => {
        setIsFocused(false)
        timeOutRef.current = setTimeout(
          setShowFocusReminder,
          FOCUS_REMINDER_TIMEOUT_MS,
          true,
        )
      }}
      onFocus={() => {
        setIsFocused(true)
        clearTimeout(timeOutRef.current)
        setShowFocusReminder(false)
      }}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
