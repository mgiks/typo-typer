import { useAppDispatch, useAppSelector } from '../../hooks'
import { resetPlayerStatus } from '../../slices/playerStatus.slice'
import { fetchText, resetTextData } from '../../slices/textData.slice'
import { resetTypingData } from '../../slices/typingData.slice'
import { resetTypingHistory } from '../../slices/typingHistory.slice'

function ResetButton() {
  const playerStartedTyping = useAppSelector((state) =>
    state.playerStatus.startedTyping
  )

  const dispatch = useAppDispatch()

  function reset() {
    dispatch(resetPlayerStatus())
    dispatch(resetTypingData())
    dispatch(resetTypingHistory())
    dispatch(resetTextData())
    dispatch(fetchText())
  }

  return (!playerStartedTyping ? <button onClick={reset}>Reset</button> : null)
}

export default ResetButton
