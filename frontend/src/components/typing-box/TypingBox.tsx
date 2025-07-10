import './TypingBox.scss'
import InputCatcher from './input-catcher/InputCatcher'

function TypingBox() {
  return (
    <div className='typing-box' data-testid='typing-box'>
      <InputCatcher />
    </div>
  )
}

export default TypingBox
