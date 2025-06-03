import { extractCorrectText, extractWrongText } from '../TextArea'
import {
  extractCorrectTextExpectationTable,
  extractWrongTextExpectationTable,
  text,
} from './fixtures'

describe('extractCorrectText', () => {
  describe(`extraction from '${text}'`, () => {
    it.each(extractCorrectTextExpectationTable)(
      `should extract $expected from wrong key at $index`,
      ({ index, expected }) => {
        expect(extractCorrectText(text, index)).toBe(expected)
      },
    )
  })
  it('should extract empty string if input is empty', () => {
    expect(extractCorrectText('', 5)).toBe('')
  })
})

describe('extractWrongText', () => {
  describe(`extraction from '${text}'`, () => {
    it.each(extractWrongTextExpectationTable)(
      `should extract $expected from wrong key at $index`,
      ({ index, expected }) => {
        expect(extractWrongText(text, index)).toBe(expected)
      },
    )
  })
  it('should extract empty string if input is empty', () => {
    expect(extractWrongText('', 5)).toBe('')
  })
})
