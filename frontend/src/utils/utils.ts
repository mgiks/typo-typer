export function calculateTypingAccuracyAndWPM(
  charsTyped: number,
  typingTimeInSeconds: number,
  correctKeyPresses: number,
  errors: number,
) {
  // A "word" is equivalent to 5 characters by convention
  function calculateWordCountForWPM(numberOfChars: number) {
    if (numberOfChars <= 0) return 0
    return Math.round(numberOfChars / 5)
  }

  function calculateTypingAccuracy(
    correctKeyPresses: number,
    errors: number,
  ) {
    if (correctKeyPresses < 0 || errors < 0) return 0
    const totalKepPresses = correctKeyPresses + errors
    return parseFloat((correctKeyPresses / totalKepPresses).toFixed(2))
  }

  // NWPM (Net Words Per Minute)
  function calculateNWPM(
    GWPM: number,
    accuracy: number,
  ) {
    if (GWPM < 0) GWPM = 0
    if (accuracy > 1) accuracy = 1
    if (accuracy < 0) accuracy = 0
    return Math.round(GWPM * accuracy)
  }

  // GWPM (Gross Words Per Minute)
  function calculateGWPM(wordCount: number, timeInMinutes: number) {
    if (wordCount <= 0 || timeInMinutes <= 0) return 0
    return Math.round(wordCount / timeInMinutes)
  }

  const typingTimeInMinutes = typingTimeInSeconds / 60
  const wordCount = calculateWordCountForWPM(charsTyped)
  const typingAccuracy = calculateTypingAccuracy(correctKeyPresses, errors)
  const GWPM = calculateGWPM(wordCount, typingTimeInMinutes)
  const NWPM = calculateNWPM(GWPM, typingAccuracy)

  return {
    GWPM: GWPM,
    NWPM: NWPM,
    typingAccuracy: typingAccuracy,
  }
}
