import './TextContainer.scss'
import Cursor from './cursor/Cursor'

function TextContainer({ text }: { text: string }) {
  const chars = text.split('')

  return (
    <div className='text-container' data-testid='text-container'>
      <Cursor />
      {chars.map((char, i) => <span key={i}>{char}</span>)}
    </div>
  )
}

export default TextContainer
