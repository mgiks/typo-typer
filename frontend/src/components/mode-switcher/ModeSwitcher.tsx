import { useState } from 'react'
import ModeModal from './mode-modal/ModeModal'

function ModeSwitcher() {
  const [showModal, setModalVisibility] = useState(false)

  return (
    <>
      <button onClick={() => setModalVisibility(true)}>
        Switch game mode
      </button>
      <ModeModal
        isVisible={showModal}
        setModalVisibility={setModalVisibility}
      />
    </>
  )
}

export default ModeSwitcher
