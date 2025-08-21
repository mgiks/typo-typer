import { useEffect, useRef, useState } from 'react'
import { useAppSelector } from '../../hooks'

function StopWatch() {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)
  const startTime = useRef(0)

  const isUserTyping = useAppSelector((state) => state.isUserTyping.value)

  useEffect(() => {
    if (!isUserTyping) return

    startTime.current = Date.now()

    const intervalId = setInterval(() => {
      setMilliSecondsElapsed(Date.now() - startTime.current)
    }, 1)

    return () => clearInterval(intervalId)
  }, [isUserTyping])

  const timer = <div role='timer'>{Math.floor(milliSecondsElapsed / 1000)}</div>

  return isUserTyping ? timer : null
}

export default StopWatch
