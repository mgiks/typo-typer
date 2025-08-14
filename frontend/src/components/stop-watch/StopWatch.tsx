type StopWatchProps = {
  visible: boolean
}

function StopWatch({ visible }: StopWatchProps) {
  const timer = <div data-testid='timer' />

  return visible ? timer : null
}

export default StopWatch
