import { calculateGWPM } from './calculateGWPM'
import { calculateNWPM } from './calculateNWPM'
import { calculateTypingAccuracy } from './calculateTypingAccuracy'
import { calculateWordCountForWPM } from './calculateWordCountForWPM'

export function calculateTypingAccuracyAndWPM(
  charsTyped: number,
  typingTimeInSeconds: number,
  correctKeyPresses: number,
  errors: number,
) {
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
