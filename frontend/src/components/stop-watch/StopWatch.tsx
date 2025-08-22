import { useEffect, useRef, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../../hooks'
import { playerStatusInitialState } from '../../slices/playerStatus.slice'
import { setTimeElapsedInMinutesTo } from '../../slices/typingStats.slice'

type StopWatchProps = {
  forceVisible?: boolean
  detachStateStore?: boolean
}

function StopWatch({ detachStateStore, forceVisible }: StopWatchProps) {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)
  const startTime = useRef(0)
  const intervalId = useRef(-1)

  const playerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)

  const playerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)

  const dispatch = detachStateStore ? () => {} : useAppDispatch()

  useEffect(() => {
    if (!playerStartedTyping && !forceVisible) return

    startTime.current = Date.now()

    intervalId.current = window.setInterval(() => {
      setMilliSecondsElapsed(Date.now() - startTime.current)
    }, 1)

    return () => clearInterval(intervalId.current)
  }, [playerStartedTyping])

  useEffect(() => {
    if (!playerFinishedTyping) return
    clearInterval(intervalId.current)

    dispatch(setTimeElapsedInMinutesTo(milliSecondsElapsed / 1000 / 60))
  }, [playerFinishedTyping, milliSecondsElapsed])

  if (playerFinishedTyping) return null

  const timer = <div role='timer'>{Math.floor(milliSecondsElapsed / 1000)}</div>

  return (forceVisible || playerStartedTyping) ? timer : null
}

export default StopWatch
