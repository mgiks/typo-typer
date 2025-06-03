import { calculateWordCount } from '../TypingStatsContainer.tsx'

const text = 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
const emptyTextWithWhitespaces = '       \t'
const textWithIrregularSpacing =
  'Lorem     ipsum  dolor     \tsit  amet,  consectetur   adipiscing    elit.'

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
