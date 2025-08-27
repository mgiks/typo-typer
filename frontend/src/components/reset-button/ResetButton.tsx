import { useAppDispatch, useAppSelector } from '../../hooks'
import { resetPlayerStatus } from '../../slices/playerStatus.slice'
import { resetTypingData } from '../../slices/typingData.slice'
import { resetTypingHistory } from '../../slices/typingHistory.slice'

function ResetButton() {
  const playerFinishedTyping = useAppSelector((state) =>
    state.playerStatus.finishedTyping
  )

  const dispatch = useAppDispatch()

  function reset() {
    dispatch(resetPlayerStatus())
    dispatch(resetTypingData())
    dispatch(resetTypingHistory())
  }

  return (playerFinishedTyping ? <button onClick={reset}>Reset</button> : null)
}

export default ResetButton
