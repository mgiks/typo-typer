import { useEffect, useRef, useState } from 'react'
import { useAppSelector } from '../../hooks'

function StopWatch() {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)
  const startTime = useRef(0)

  const hasUserStartedTyping = useAppSelector((state) =>
    state.playerStatus.startedTyping
  )

  useEffect(() => {
    if (!hasUserStartedTyping) return

    startTime.current = Date.now()

    const intervalId = setInterval(() => {
      setMilliSecondsElapsed(Date.now() - startTime.current)
    }, 1)

    return () => clearInterval(intervalId)
  }, [hasUserStartedTyping])

  const timer = <div role='timer'>{Math.floor(milliSecondsElapsed / 1000)}</div>

  return hasUserStartedTyping ? timer : null
}

export default StopWatch
