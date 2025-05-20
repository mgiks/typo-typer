import { parseText } from '../parseText'
import {
  text,
  textWithIrregularSpacing,
  textWithNewLineChars,
} from './fixtures'

describe('parseText', () => {
  it('should remove newline chars', () => {
    expect(parseText(textWithNewLineChars)).toBe(text)
  })
  it('should squash multiple spaces into one', () => {
    expect(parseText(textWithIrregularSpacing)).toBe(text)
  })
})
