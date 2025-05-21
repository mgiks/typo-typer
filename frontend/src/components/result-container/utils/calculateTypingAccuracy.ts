export function calculateTypingAccuracy(
  correctKeyPresses: number,
  errors: number,
) {
  if (correctKeyPresses < 0 || errors < 0) return 0
  const totalKepPresses = correctKeyPresses + errors
  return parseFloat((correctKeyPresses / totalKepPresses).toFixed(2))
}
