import { useEffect, useRef, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../../hooks'
import { playerStatusInitialState } from '../../slices/playerStatus.slice'
import {
  setTimeElapsedInMinutesTo,
  typingStatsInitialState,
} from '../../slices/typingStats.slice'
import { addResultGraphPoint } from '../../slices/resultGraph.slice'
import {
  calculateAccuracy,
  calculateAdjustedWpm,
  calculateRawWpm,
} from '../../utils/typing-results.utils'

const TIME_FRAME_IN_MS = 100

type StopWatchProps = {
  forceVisible?: boolean
  detachStateStore?: boolean
}

function StopWatch({ detachStateStore, forceVisible }: StopWatchProps) {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)
  const startTime = useRef(0)
  const currentTime = useRef(0)
  const intervalId = useRef(-1)

  const playerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)

  const playerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)

  const totalKeysPressed = detachStateStore
    ? typingStatsInitialState.totalKeysPressed
    : useAppSelector((state) => state.typingStats.totalKeysPressed)

  const correctKeysPressed = detachStateStore
    ? typingStatsInitialState.correctKeysPressed
    : useAppSelector((state) => state.typingStats.correctKeysPressed)

  const dispatch = detachStateStore ? () => {} : useAppDispatch()

  useEffect(() => {
    if (!playerStartedTyping && !forceVisible) return

    startTime.current = Date.now()
    currentTime.current = startTime.current

    intervalId.current = window.setInterval(() => {
      currentTime.current += TIME_FRAME_IN_MS
      const timeDiff = currentTime.current - startTime.current
      setMilliSecondsElapsed(timeDiff)
    }, TIME_FRAME_IN_MS)

    return () => clearInterval(intervalId.current)
  }, [playerStartedTyping])

  const prevTimePoint = useRef(-1)

  useEffect(() => {
    if (milliSecondsElapsed && prevTimePoint.current !== milliSecondsElapsed) {
      const acc = calculateAccuracy(totalKeysPressed, correctKeysPressed)
      const rawWpm = calculateRawWpm(
        milliSecondsElapsed / 1000 / 60,
        totalKeysPressed,
      )
      const adjustedWpm = calculateAdjustedWpm(rawWpm, acc)

      dispatch(addResultGraphPoint({
        time: milliSecondsElapsed,
        acc: acc,
        wpm: adjustedWpm,
        errs: totalKeysPressed - correctKeysPressed,
      }))
    }
    prevTimePoint.current = milliSecondsElapsed
  }, [totalKeysPressed, correctKeysPressed, milliSecondsElapsed])

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
