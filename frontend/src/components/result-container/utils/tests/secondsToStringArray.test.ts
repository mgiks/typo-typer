import { stringArrayFrom1to10 } from './fixtures'
import { secondsToStringArray } from '../secondsToStringArray'

describe('secondsToStringArray', () => {
  it('should return string array from 1 to 10 from 10', () => {
    expect(secondsToStringArray(10)).toEqual(stringArrayFrom1to10)
  })
  it('should return empty array from 0', () => {
    expect(secondsToStringArray(0)).toEqual([])
  })
  it('should return empty array from -100', () => {
    expect(secondsToStringArray(-100)).toEqual([])
  })
})
