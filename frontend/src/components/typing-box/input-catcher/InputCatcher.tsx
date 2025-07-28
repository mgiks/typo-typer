import { useEffect, useState } from 'react'

type InputCatcherProps = {
  text: string
  lastTypedLetterIndexSetter: (i: number) => void
  incorrectLetterIndexSetter: (i: number) => void
}

function InputCatcher(
  { text, lastTypedLetterIndexSetter, incorrectLetterIndexSetter }:
    InputCatcherProps,
) {
  const [input, setInput] = useState('')

  useEffect(() => {
    if (input.length == 0) {
      return
    }

    const inputLastIndex = input.length - 1
    lastTypedLetterIndexSetter(inputLastIndex)

    if (input.length > 0 && (text[inputLastIndex] !== input.at(-1))) {
      incorrectLetterIndexSetter(inputLastIndex)
    }
  }, [input])

  return (
    <textarea
      className='typing-box__input-catcher'
      data-testid='input-catcher'
      onInput={(event) => setInput(event.currentTarget.value)}
      autoFocus
    >
    </textarea>
  )
}

export default InputCatcher
