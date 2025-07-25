import './TextContainer.scss'
import Cursor from './cursor/Cursor'

type TextContainerProps = {
  text: string
  lastTypedLetterIndex: number
  incorrectTextStartIndex: number
}

function TextContainer(
  { text, lastTypedLetterIndex, incorrectTextStartIndex }: TextContainerProps,
) {
  let correctText = ''
  let incorrectText = ''

  if (incorrectTextStartIndex == -1) {
    correctText = text.slice(0, lastTypedLetterIndex)
  } else {
    correctText = text.slice(0, incorrectTextStartIndex)
    incorrectText = text.slice(
      incorrectTextStartIndex,
      lastTypedLetterIndex + 1,
    )
  }

  return (
    <div className='text-container' data-testid='text-container'>
      <Cursor />
      <span data-testid='correct-text'>{correctText}</span>
      <span data-testid='incorrect-text'>{incorrectText}</span>
      <span>{text.slice(lastTypedLetterIndex + 1)}</span>
    </div>
  )
}

export default TextContainer
