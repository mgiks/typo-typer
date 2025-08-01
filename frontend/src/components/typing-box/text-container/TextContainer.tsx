import './TextContainer.scss'
import Cursor from './cursor/Cursor'

type TextContainerProps = {
  text: string
  lastTypedLetterIndex: number
  incorrectTextStartIndex: number
  focusSetter: (i: boolean) => void
  refToFocus: {
    current: {
      focus: () => void
    } | null
  }
}

function TextContainer(
  {
    text,
    lastTypedLetterIndex,
    incorrectTextStartIndex,
    focusSetter,
    refToFocus,
  }: TextContainerProps,
) {
  let correctText = ''
  let incorrectText = ''

  if (incorrectTextStartIndex == -1) {
    correctText = text.slice(0, lastTypedLetterIndex + 1)
  } else {
    correctText = text.slice(0, incorrectTextStartIndex)
    incorrectText = text.slice(
      incorrectTextStartIndex,
      lastTypedLetterIndex + 1,
    )
  }

  return (
    <div
      className='text-container'
      data-testid='text-container'
      onClick={() => (focusSetter(true), refToFocus.current?.focus())}
    >
      <span className='text-container__text_correct' data-testid='correct-text'>
        {correctText}
      </span>
      <span
        className='text-container__text_incorrect'
        data-testid='incorrect-text'
      >
        {incorrectText}
      </span>
      <Cursor />
      <span>{text.slice(lastTypedLetterIndex + 1)}</span>
    </div>
  )
}

export default TextContainer
