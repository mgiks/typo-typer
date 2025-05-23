import { useEffect, useState } from 'react'
import './InactivityCurtain.css'

function InactivityCurtain() {
  // 'typing-area' is needed so that the curtain doesn't show initially
  const [activeElementId, setActiveElementId] = useState('')

  useEffect(() => {
    const updateActiveElement = () => {
      const activeElement = document.activeElement
      activeElement && setActiveElementId(activeElement.id)
    }

    const events = ['click', 'keypress', 'focusin']
    events.forEach((event) =>
      document.addEventListener(event, updateActiveElement)
    )
  }, [])

  const inactivityCurtain = (
    <div id='inactivity-curtain'>
      Click here or type any key to continue
    </div>
  )

  return activeElementId !== 'typing-area' ? inactivityCurtain : null
}

export default InactivityCurtain
