import { calculateTypingAccuracyAndWPM } from '../ResultContainer'

describe('calculateTypingAccuracyAndWPM', () => {
  it('should calculate GWPM, NWPM and typingAccuracy', () => {
    calculateTypingAccuracyAndWPM(
      charsTyped,
      typingTimeInSeconds,
      correctKeypresses,
      errors,
    )

    expect(result).toEqual({
      GWPM: GWPM,
      NWPM: NWPM,
      typingAccuracy: typingAccuracy,
    })
  })
})
