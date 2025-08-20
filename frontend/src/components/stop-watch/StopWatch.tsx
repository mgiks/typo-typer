import { useEffect, useState } from 'react'
import { useAppSelector } from '../../hooks'

function StopWatch() {
  const [secondsElapsed, setSecondsElapsed] = useState(0)

  const isUserTyping = useAppSelector((state) => state.isUserTyping.value)

  useEffect(() => {
    if (!isUserTyping) return

    const intervalId = setInterval(() => {
      setSecondsElapsed((s) => s + 1)
    }, 1000)

    return () => clearInterval(intervalId)
  }, [isUserTyping])

  const timer = <div role='timer'>{secondsElapsed}</div>

  return isUserTyping ? timer : null
}

export default StopWatch
