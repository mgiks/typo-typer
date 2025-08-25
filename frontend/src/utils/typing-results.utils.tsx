export function calculateAccuracy(
  totalKeysPressed: number,
  correctKeysPressed: number,
) {
  if (!totalKeysPressed) return 0
  // Rounding until the nearest hundredth of a decimanl
  return Math.ceil(correctKeysPressed / totalKeysPressed * 100) / 100
}

export function calculateRawWpm(
  timeElapsedInMinutes: number,
  totalKeysPressed: number,
) {
  return timeElapsedInMinutes
    ? Math.floor(totalKeysPressed / 5 / timeElapsedInMinutes)
    : 0
}

export function calculateAdjustedWpm(rawWpm: number, acc: number) {
  return Math.ceil(rawWpm * acc * 100) / 100
}
