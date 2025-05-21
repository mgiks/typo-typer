import { calculateGWPM } from '../calculateGWPM'

describe('calculateGWPM', () => {
  it('should calculate 10 GWPM from 10 words and 1 minute', () => {
    expect(calculateGWPM(10, 1)).toBe(10)
  })
  it('should calculate 40 GWPM from 20 words and 0.5 minutes', () => {
    expect(calculateGWPM(20, 0.5)).toBe(40)
  })
  it('should calculate 0 GWPM from 0 words and 100 minutes', () => {
    expect(calculateGWPM(0, 100)).toBe(0)
  })
  it('should calculate 0 GWPM from -10 words and 1 minute', () => {
    expect(calculateGWPM(-10, 1)).toBe(0)
  })
  it('should calculate 0 GWPM from 10 words and -1 minute', () => {
    expect(calculateGWPM(10, -1)).toBe(0)
  })
  it('should calculate 0 GWPM from -10 words and -100 minute', () => {
    expect(calculateGWPM(-10, -100)).toBe(0)
  })
})
