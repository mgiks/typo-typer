import React, { useRef } from 'react'

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
  const lastIncorrectLetterIndex = useRef(-1)

  const handleInput = (event: React.FormEvent<HTMLTextAreaElement>) => {
    const input = event.currentTarget.value

    const inputLastIndex = input.length - 1
    setLastTypedIndex(inputLastIndex)

    if (inputLastIndex === -1) {
      lastIncorrectLetterIndex.current = -1
      setIncorrectTextStartIndex(-1)
      return
    }

    if (text[inputLastIndex] !== input.at(-1)) {
      if (
        lastIncorrectLetterIndex.current === -1 ||
        inputLastIndex < lastIncorrectLetterIndex.current
      ) {
        setIncorrectTextStartIndex(inputLastIndex)
        lastIncorrectLetterIndex.current = inputLastIndex
      }
    } else if (inputLastIndex < lastIncorrectLetterIndex.current) {
      setIncorrectTextStartIndex(-1)
      lastIncorrectLetterIndex.current = -1
    }
  }

  return (
    <textarea
      ref={ref}
      className='typing-box__input-catcher'
      onInput={handleInput}
      onBlur={() => {
        setIsFocused(false)
        setShowFocusReminder(true)
      }}
      onFocus={() => {
        setIsFocused(true)
        setShowFocusReminder(false)
      }}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
