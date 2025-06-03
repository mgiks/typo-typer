import {
  isControlKeyExpectationTable,
  text,
  textWithIrregularSpacing,
  textWithNewLineChars,
} from './fixtures'
import { isControlKey, parseText } from '../TypingArea'

describe('parseText', () => {
  it('should remove newline chars', () => {
    expect(parseText(textWithNewLineChars)).toBe(text)
  })
  it('should squash multiple spaces into one', () => {
    expect(parseText(textWithIrregularSpacing)).toBe(text)
  })
})

describe('isControlKey', () => {
  it.each(isControlKeyExpectationTable)(
    `should return $expected from $key`,
    ({ key, expected }) => {
      expect(isControlKey(key)).toBe(expected)
    },
  )
})
