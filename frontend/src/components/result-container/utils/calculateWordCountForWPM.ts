// A "word" is equivalent to 5 characters by convention
export function calculateWordCountForWPM(numberOfChars: number) {
  if (numberOfChars <= 0) return 0
  return Math.round(numberOfChars / 5)
}
