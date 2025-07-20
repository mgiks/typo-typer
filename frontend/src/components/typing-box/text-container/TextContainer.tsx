import './TextContainer.scss'
import Cursor from './cursor/Cursor'

function TextContainer({ text }: { text: string }) {
  return (
    <div data-testid='text-container'>
      <Cursor />
      {text}
    </div>
  )
}

export default TextContainer
