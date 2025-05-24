import './TypingContainer.css'
import TextArea from './TextArea'
import TypingArea from './TypingArea'
import { focusElement } from './utils/focusElement'
import { useEffect, useRef } from 'react'
import { useIsDoneTyping } from '../../stores/TypingStatsStore'

function TypingContainer(
  { shouldLogInFormBeShown }: { shouldLogInFormBeShown: boolean },
) {
  const typingContainerRef = useRef<HTMLDivElement>(null)
  const typingAreaRef = useRef<HTMLTextAreaElement | null>(null)

  const isDoneTyping = useIsDoneTyping()

  function focusTypingArea() {
    typingAreaRef.current && focusElement(typingAreaRef.current)
  }

  const areOutsideInputsNeeded = useRef(false)
  useEffect(() => {
    areOutsideInputsNeeded.current = shouldLogInFormBeShown
  }, [shouldLogInFormBeShown])

  useEffect(() => {
    function handleKeyDown() {
      if (areOutsideInputsNeeded.current) return

      focusTypingArea()
    }

    document.addEventListener('keypress', handleKeyDown)
    return () => document.removeEventListener('keypress', handleKeyDown)
  }, [])

  useEffect(() => {
    focusTypingArea()
  }, [isDoneTyping])

  return (
    <div
      id='typing-container'
      ref={typingContainerRef}
      onClick={focusTypingArea}
    >
      <TypingArea ref={typingAreaRef} />
      <TextArea typingContainerRef={typingContainerRef} />
    </div>
  )
}

export default TypingContainer
