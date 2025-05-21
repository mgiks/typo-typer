// GWPM (Gross Words Per Minute)
export function calculateGWPM(wordCount: number, timeInMinutes: number) {
  if (wordCount <= 0 || timeInMinutes <= 0) return 0
  return Math.round(wordCount / timeInMinutes)
}
