import { createPortal } from 'react-dom'
import './ModeModal.scss'

type Setter<T> = (i: T) => void

type ModeModalProps = {
  isVisible: boolean
  setModalVisibility: Setter<boolean>
}

function ModeModal({ isVisible, setModalVisibility }: ModeModalProps) {
  if (!isVisible) return null

  return createPortal(
    <div
      className='mode-modal-container'
      onClick={() => setModalVisibility(false)}
    >
      <div className='mode-modal' onClick={(event) => event.stopPropagation()}>
        <div className='mode-modal__header'>
          <label>
            Choose mode:
          </label>
          <button
            className='mode-modal__close-button'
            onClick={() => {
              setModalVisibility(false)
            }}
          >
            x
          </button>
        </div>
        <div className='mode-modal__option-section'>
          <button
            className='mode-modal__mode-button'
            onClick={() => {
              setModalVisibility(false)
            }}
          >
            Solo
          </button>
          <button
            className='mode-modal__mode-button'
            onClick={() => {
              setModalVisibility(false)
            }}
          >
            Multiplayer
          </button>
        </div>
      </div>
    </div>,
    document.getElementById('modal-container')!,
  )
}

export default ModeModal
