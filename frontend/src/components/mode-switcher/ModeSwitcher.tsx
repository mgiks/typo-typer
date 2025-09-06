import { useState } from 'react'
import ModeModal from './mode-modal/ModeModal'
import './ModeSwitcher.scss'
import { useAppSelector } from '../../hooks'

function ModeSwitcher() {
  const [showModal, setModalVisibility] = useState(false)
  const playerMode = useAppSelector((state) => state.playerMode.mode)

  return (
    <>
      <button
        className='mode-switcher__button'
        onClick={() => setModalVisibility(true)}
      >
        {playerMode}
      </button>
      <ModeModal
        isVisible={showModal}
        setModalVisibility={setModalVisibility}
      />
    </>
  )
}

export default ModeSwitcher
