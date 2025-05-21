export function secondsToStringArray(seconds: number) {
  if (seconds <= 0) return []
  return Array(seconds).fill(0).map((_, i) => (seconds - i).toString())
    .reverse()
}
