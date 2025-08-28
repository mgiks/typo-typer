import { useEffect, useRef, useState } from 'react'
import { useAppDispatch, useAppSelector } from '../../hooks'
import { typingDataInitialState } from '../../slices/typingData.slice'
import {
  addTypingHistoryPoint,
  setLastRecordedMomentTo,
} from '../../slices/typingHistory.slice'
import {
  calculateAccuracy,
  calculateAdjustedWpm,
  calculateRawWpm,
} from '../../utils/typing-results.utils'
import { playerStatusInitialState } from '../../slices/playerStatus.slice'

const TIME_FRAME_IN_MS = 100

type StopWatchProps = { forceVisible?: boolean; detachStateStore?: boolean }

function StopWatch({ detachStateStore, forceVisible }: StopWatchProps) {
  const [milliSecondsElapsed, setMilliSecondsElapsed] = useState(0)

  const startTime = useRef(0)
  const currentTime = useRef(0)
  const intervalId = useRef(-1)

  // Keeps interval aware of these values
  const totalKeysPressed = useRef(0)
  const correctKeysPressed = useRef(0)

  const playerStartedTyping = detachStateStore
    ? playerStatusInitialState.startedTyping
    : useAppSelector((state) => state.playerStatus.startedTyping)

  const playerFinishedTyping = detachStateStore
    ? playerStatusInitialState.finishedTyping
    : useAppSelector((state) => state.playerStatus.finishedTyping)

  const totalKeysPressedState = detachStateStore
    ? typingDataInitialState.totalKeysPressed
    : useAppSelector((state) => state.typingData.totalKeysPressed)

  const correctKeysPressedState = detachStateStore
    ? typingDataInitialState.correctKeysPressed
    : useAppSelector((state) => state.typingData.correctKeysPressed)

  const dispatch = detachStateStore ? () => {} : useAppDispatch()

  useEffect(() => {
    totalKeysPressed.current = totalKeysPressedState
  }, [totalKeysPressedState])

  useEffect(() => {
    correctKeysPressed.current = correctKeysPressedState
  }, [correctKeysPressedState])

  const prevTimePoint = useRef(-1)

  useEffect(() => {
    if (!playerStartedTyping && !forceVisible) return

    startTime.current = Date.now()
    currentTime.current = startTime.current

    intervalId.current = window.setInterval(() => {
      currentTime.current += TIME_FRAME_IN_MS
      const timeDiffInMs = currentTime.current - startTime.current

      setMilliSecondsElapsed(timeDiffInMs)

      if (timeDiffInMs) {
        const timeDiffInSeconds = timeDiffInMs / 1000
        const acc = calculateAccuracy(
          totalKeysPressed.current,
          correctKeysPressed.current,
        )
        const rawWpm = calculateRawWpm(
          timeDiffInSeconds / 60,
          totalKeysPressed.current,
        )
        const adjustedWpm = calculateAdjustedWpm(rawWpm, acc)

        dispatch(addTypingHistoryPoint({
          timeInSeconds: timeDiffInSeconds,
          acc: acc,
          wpm: adjustedWpm,
          errs: totalKeysPressed.current - correctKeysPressed.current,
        }))
      }

      prevTimePoint.current = milliSecondsElapsed
    }, TIME_FRAME_IN_MS)

    return () => clearInterval(intervalId.current)
  }, [playerStartedTyping])

  useEffect(() => {
    if (!playerFinishedTyping) return
    clearInterval(intervalId.current)
    const finaltimeDiffInMs = currentTime.current - startTime.current

    const finaltimeDiffInSeconds = finaltimeDiffInMs / 1000
    const finalAcc = calculateAccuracy(
      totalKeysPressed.current,
      correctKeysPressed.current,
    )
    const finalRawWpm = calculateRawWpm(
      finaltimeDiffInSeconds / 60,
      totalKeysPressed.current,
    )
    const finalAdjustedWpm = calculateAdjustedWpm(finalRawWpm, finalAcc)

    dispatch(setLastRecordedMomentTo({
      acc: finalAcc,
      wpm: finalAdjustedWpm,
      errs: totalKeysPressed.current - correctKeysPressed.current,
    }))
  }, [playerFinishedTyping, milliSecondsElapsed])

  if (playerFinishedTyping) return null

  const timer = <div role='timer'>{Math.floor(milliSecondsElapsed / 1000)}</div>

  return (forceVisible || playerStartedTyping) ? timer : null
}

export default StopWatch
