import { useEffect, useRef, useState } from 'react'
import { useAppSelector } from '../../hooks'
import { playerStatusInitialState } from '../../slices/playerStatus.slice'

type StopWatchProps = {
  forceVisible?: boolean
  detachStateStore?: boolean
}

function StopWatch({ detachStateStore, forceVisible }: StopWatchProps) {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)
  const startTime = useRef(0)

  const playerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)

  const playerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)

  useEffect(() => {
    if (!playerStartedTyping && !forceVisible) return

    startTime.current = Date.now()

    const intervalId = setInterval(() => {
      setMilliSecondsElapsed(Date.now() - startTime.current)
    }, 1)

    return () => clearInterval(intervalId)
  }, [playerStartedTyping])

  if (playerFinishedTyping) return null

  const timer = <div role='timer'>{Math.floor(milliSecondsElapsed / 1000)}</div>

  return (forceVisible || playerStartedTyping) ? timer : null
}

export default StopWatch
