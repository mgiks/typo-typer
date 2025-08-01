import { useEffect, useState } from 'react'

type InputCatcherProps = {
  text: string
  lastTypedLetterIndexSetter: (i: number) => void
  incorrectTextStartIndexSetter: (i: number) => void
  focusedSetter: (i: boolean) => void
}

function InputCatcher(
  {
    text,
    lastTypedLetterIndexSetter,
    incorrectTextStartIndexSetter,
    focusedSetter,
  }: InputCatcherProps,
) {
  const [input, setInput] = useState('')
  const [lastIncorrectLetterIndex, setLastIncorrectLetterIndex] = useState(-1)

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
      className='typing-box__input-catcher'
      data-testid='input-catcher'
      onInput={(event) => setInput(event.currentTarget.value)}
      onBlur={() => focusedSetter(false)}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
