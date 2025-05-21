import { calculateTypingAccuracyAndWPM } from '../calculateTypingAccuracyAndWPM'
import * as accuracyModule from '../calculateTypingAccuracy.tsx'
import * as gwpmModule from '../calculateGWPM.tsx'
import * as nwpmModule from '../calculateNWPM.tsx'
import * as wordCountModule from '../calculateWordCountForWPM.tsx'

describe('calculateTypingAccuracyAndWPM', () => {
  it('should calculate GWPM, NWPM and typingAccuracy correctly', () => {
    const charsTyped = 300
    const typingTimeInSeconds = 60
    const correctKeypresses = 100
    const errors = 5

    const wordCount = 100
    const typingTimeInMinutes = 1
    const typingAccuracy = 0.5
    const GWPM = 100
    const NWPM = 50

    vi.spyOn(wordCountModule, 'calculateWordCountForWPM').mockReturnValue(
      wordCount,
    )
    vi.spyOn(accuracyModule, 'calculateTypingAccuracy').mockReturnValue(
      typingAccuracy,
    )
    vi.spyOn(gwpmModule, 'calculateGWPM').mockReturnValue(GWPM)
    vi.spyOn(nwpmModule, 'calculateNWPM').mockReturnValue(NWPM)

    const result = calculateTypingAccuracyAndWPM(
      charsTyped,
      typingTimeInSeconds,
      correctKeypresses,
      errors,
    )

    expect(wordCountModule.calculateWordCountForWPM).toHaveBeenCalledWith(
      charsTyped,
    )
    expect(accuracyModule.calculateTypingAccuracy).toHaveBeenCalledWith(
      correctKeypresses,
      errors,
    )
    expect(gwpmModule.calculateGWPM).toHaveBeenCalledWith(
      wordCount,
      typingTimeInMinutes,
    )
    expect(nwpmModule.calculateNWPM).toHaveBeenCalledWith(GWPM, typingAccuracy)

    expect(result).toEqual({
      GWPM: GWPM,
      NWPM: NWPM,
      typingAccuracy: typingAccuracy,
    })
  })
})
