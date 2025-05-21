import { isControlKeyExpectationTable } from './fixtures'
import { isControlKey } from '../isControlKey'

describe('isControlKey', () => {
  it.each(isControlKeyExpectationTable)(
    `should return $expected from $key`,
    ({ key, expected }) => {
      expect(isControlKey(key)).toBe(expected)
    },
  )
})
