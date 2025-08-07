import React, { useEffect, useRef, useState } from 'react'

export const SHOW_REMINDER_TO_FOCUS_DELAY = 750

type Setter<T> = (i: T) => void

type InputCatcherProps = {
  text: string
  setIsFocused: Setter<boolean>
  setLastTypedIndex: Setter<number>
  setShowReminderToFocus: Setter<boolean>
  setIncorrectTextStartIndex: Setter<number>
  ref: React.Ref<HTMLTextAreaElement | null>
}

function InputCatcher(
  {
    ref,
    text,
    setIsFocused,
    setLastTypedIndex,
    setShowReminderToFocus,
    setIncorrectTextStartIndex,
  }: InputCatcherProps,
) {
  const [input, setInput] = useState('')
  const [lastIncorrectLetterIndex, setLastIncorrectLetterIndex] = useState(-1)
  const timeOutRef = useRef(-1)

  useEffect(() => {
    const inputLastIndex = input.length - 1
    setLastTypedIndex(inputLastIndex)

    if (inputLastIndex == -1) {
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
      data-testid='input-catcher'
      onInput={(event) => setInput(event.currentTarget.value)}
      onBlur={() => {
        setIsFocused(false)
        timeOutRef.current = setTimeout(
          setShowReminderToFocus,
          SHOW_REMINDER_TO_FOCUS_DELAY,
          true,
        )
      }}
      onFocus={() => {
        setIsFocused(true)
        clearTimeout(timeOutRef.current)
        setShowReminderToFocus(false)
      }}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
