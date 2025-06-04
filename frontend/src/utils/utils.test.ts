import { calculateTypingAccuracyAndWPM } from './utils'

describe('calculateTypingAccuracyAndWPM', () => {
  it('should calculate GWPM, NWPM and typingAccuracy', () => {
    const charsTyped = 100
    const typingTimeInSeconds = 50
    const correctKeypresses = 50
    const errors = 50

    const result = calculateTypingAccuracyAndWPM(
      charsTyped,
      typingTimeInSeconds,
      correctKeypresses,
      errors,
    )

    expect(result).toEqual({
      GWPM: 24,
      NWPM: 12,
      typingAccuracy: 0.5,
    })
  })
})
