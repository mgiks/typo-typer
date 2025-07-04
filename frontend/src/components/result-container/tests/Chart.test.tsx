import { secondsToStringArray } from '../Chart.tsx'

describe('secondsToStringArray', () => {
  it('should return string array from 1 to 10 from 10', () => {
    const stringArrayFrom1to10 = [
      '1',
      '2',
      '3',
      '4',
      '5',
      '6',
      '7',
      '8',
      '9',
      '10',
    ]

    expect(secondsToStringArray(10)).toEqual(stringArrayFrom1to10)
  })

  it('should return empty array from 0', () => {
    expect(secondsToStringArray(0)).toEqual([])
  })

  it('should return empty array from -100', () => {
    expect(secondsToStringArray(-100)).toEqual([])
  })
})
