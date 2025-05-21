import { calculateWordCountForWPM } from '../calculateWordCountForWPM'

describe('calculateWordCountForWPM', () => {
  it('should count 50 words from 250 ', () => {
    expect(calculateWordCountForWPM(250)).toEqual(50)
  })
  it('should count 0 words from 0 ', () => {
    expect(calculateWordCountForWPM(0)).toEqual(0)
  })
  it('should count 0 words from -100 ', () => {
    expect(calculateWordCountForWPM(-100)).toEqual(0)
  })
})
