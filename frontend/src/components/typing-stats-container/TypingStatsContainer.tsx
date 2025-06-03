import { useEffect, useRef, useState } from 'react'
import './TypingStatsContainer.css'
import { useCorrectText, useCursorIndex, useText } from '../../stores/TextStore'
import {
  useCorrectKeyCount,
  useCursorMoved,
  useIsDoneTyping,
  useIsStopWatchRunning,
  useTimeElapsed,
  useTypingStatsActions,
  useWrongKeyCount,
} from '../../stores/TypingStatsStore'
import PlayerModeSwitcher from './PlayerModeSwitcher'
import { Result, useResultActions } from '../../stores/ResultStore'
import { calculateTypingAccuracyAndWPM } from '../result-container/utils/calculateTypingAccuracyAndWPM'

function TypingStatsContainer() {
  const cursorMoved = useCursorMoved()
  const isDoneTyping = useIsDoneTyping()
  const text = useText()
  const errors = useWrongKeyCount()
  const correctText = useCorrectText()
  const correctKeyPresses = useCorrectKeyCount()
  const isStopwatchRunning = useIsStopWatchRunning()
  const timeElapsed = useTimeElapsed()
  const cursorIndex = useCursorIndex()

  const { addResult } = useResultActions()

  const { setTypingTime, setTimeElapsed, startStopwatch, stopStopwatch } =
    useTypingStatsActions()

  const [typingStartTime, setTypingStartTime] = useState(0)

  useEffect(() => {
    if (!isDoneTyping) return

    stopStopwatch()
    setTypingTime(timeElapsed)
  }, [isDoneTyping])

  useEffect(() => {
    if (!cursorMoved) return

    startStopwatch()
    setTypingStartTime(Date.now())
  }, [cursorMoved])

  const stopWarchRef = useRef<NodeJS.Timeout>(undefined)
  useEffect(() => {
    if (!isStopwatchRunning) {
      clearTimeout(stopWarchRef.current)
      return
    }

    stopWarchRef.current = setInterval(() => {
      setTimeElapsed(Math.round((Date.now() - typingStartTime) / 1000))
    }, 1000)
  }, [isStopwatchRunning])

  useEffect(() => {
    if (!timeElapsed) return

    const { GWPM, NWPM, typingAccuracy } = calculateTypingAccuracyAndWPM(
      cursorIndex,
      timeElapsed,
      correctKeyPresses,
      errors,
    )

    const result: Result = {
      GWPM: GWPM,
      NWPM: NWPM,
      typingAccuracy: typingAccuracy,
      time: timeElapsed,
      errors: errors,
    }

    addResult(result)
  }, [timeElapsed, isDoneTyping])

  return (
    <div id='typing-stats-container'>
      <PlayerModeSwitcher />
      <div id='stopwatch'>{timeElapsed}</div>
      <div id='word-count'>
        {calculateWordCount(correctText)} / {calculateWordCount(text)}
      </div>
    </div>
  )
}

export function calculateWordCount(text: string) {
  text = text.trim()
  text = text.replaceAll(/\s+/g, ' ')
  if (!text) return 0
  return text.split(' ').length
}

export default TypingStatsContainer
