import './TextContainer.scss'
import Cursor from './cursor/Cursor'

export type TextContainerProps = {
  showCursor: boolean
  text: string
  lastTypedIndex: number
  incorrectTextStartIndex: number
  cursorRef: React.Ref<HTMLSpanElement | null>
}

function TextContainer(
  {
    text,
    showCursor,
    lastTypedIndex,
    incorrectTextStartIndex,
    cursorRef,
  }: TextContainerProps,
) {
  let correctText = ''
  let incorrectText = ''

  if (incorrectTextStartIndex === -1) {
    correctText = text.slice(0, lastTypedIndex + 1)
  } else {
    correctText = text.slice(0, incorrectTextStartIndex)
    incorrectText = text.slice(
      incorrectTextStartIndex,
      lastTypedIndex + 1,
    )
  }

  return (
    <div
      className='text-container'
      role='status'
    >
      <span className='text-container__text_correct' aria-label='Correct text'>
        {correctText}
      </span>
      <span
        className='text-container__text_incorrect'
        aria-label='Incorrect text'
      >
        {incorrectText}
      </span>
      <Cursor
        ref={cursorRef}
        visible={showCursor}
      />
      <span>{text.slice(lastTypedIndex + 1)}</span>
    </div>
  )
}

export default TextContainer
