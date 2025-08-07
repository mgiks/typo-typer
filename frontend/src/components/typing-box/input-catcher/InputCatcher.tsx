import React, { useEffect, useRef, useState } from 'react'

export const BLUR_TIMEOUT = 750

type InputCatcherProps = {
  text: string
  lastTypedLetterIndexSetter: (i: number) => void
  incorrectTextStartIndexSetter: (i: number) => void
  focusSetter: (i: boolean) => void
  ref: React.Ref<HTMLTextAreaElement | null>
}

function InputCatcher(
  {
    text,
    lastTypedLetterIndexSetter,
    incorrectTextStartIndexSetter,
    focusSetter,
    ref,
  }: InputCatcherProps,
) {
  const [input, setInput] = useState('')
  const [lastIncorrectLetterIndex, setLastIncorrectLetterIndex] = useState(-1)
  const timeOutRef = useRef(-1)

  useEffect(() => {
    const inputLastIndex = input.length - 1
    lastTypedLetterIndexSetter(inputLastIndex)

    if (inputLastIndex == -1) {
      setLastIncorrectLetterIndex(-1)
      incorrectTextStartIndexSetter(-1)
      return
    }

    if (text[inputLastIndex] !== input.at(-1)) {
      if (
        lastIncorrectLetterIndex === -1 ||
        inputLastIndex < lastIncorrectLetterIndex
      ) {
        incorrectTextStartIndexSetter(inputLastIndex)
        setLastIncorrectLetterIndex(inputLastIndex)
      }
    } else if (inputLastIndex < lastIncorrectLetterIndex) {
      incorrectTextStartIndexSetter(-1)
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
        timeOutRef.current = setTimeout(focusSetter, BLUR_TIMEOUT, false)
      }}
      onFocus={() => {
        clearTimeout(timeOutRef.current)
        focusSetter(true)
      }}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
