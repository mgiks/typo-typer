import { useAppSelector } from '../../hooks'

function StopWatch() {
  const timer = <div role='timer' />

  const isUserTyping = useAppSelector((state) => state.isUserTyping.value)

  return isUserTyping ? timer : null
}

export default StopWatch
