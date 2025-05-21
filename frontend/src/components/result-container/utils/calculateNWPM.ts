// NWPM (Net Words Per Minute)
export function calculateNWPM(
  GWPM: number,
  accuracy: number,
) {
  if (GWPM < 0) GWPM = 0
  if (accuracy > 1) accuracy = 1
  if (accuracy < 0) accuracy = 0
  return Math.round(GWPM * accuracy)
}
