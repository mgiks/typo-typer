import {
  emptyTextWithWhitespaces,
  text,
  textWithIrregularSpacing,
} from './fixtures'
import { calculateWordCount } from '../calculateWordCount'

describe('calculateWordCount', () => {
  it('should count 8 words from regular text', () =>
    expect(calculateWordCount(text)).toBe(8))
  it('should count 0 words from empty text', () =>
    expect(calculateWordCount('')).toBe(0))
  it('should count 0 words from text with spaces only', () =>
    expect(calculateWordCount(emptyTextWithWhitespaces)).toBe(0))
  it('should count 8 words from text with irregular spacing', () =>
    expect(calculateWordCount(textWithIrregularSpacing)).toBe(8))
})
