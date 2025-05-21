import { calculateNWPM } from '../calculateNWPM'

describe('calculateNWPM', () => {
  it('should calculate 100 NWPM from 100 GWPM and 100% accuracy', () => {
    expect(calculateNWPM(100, 1)).toBe(100)
  })
  it('should calculate 50 NWPM from 100 GWPM and 50% accuracy', () => {
    expect(calculateNWPM(100, 0.5)).toBe(50)
  })
  it('should calculate 0 NWPM from 100 GWPM and 0% accuracy', () => {
    expect(calculateNWPM(100, 0)).toBe(0)
  })
  it('should calculate 100 NWPM from 100 GWPM and 200% accuracy', () => {
    expect(calculateNWPM(100, 2)).toBe(100)
  })
  it('should calculate 0 NWPM from 100 GWPM and -1% accuracy', () => {
    expect(calculateNWPM(100, -0.01)).toBe(0)
  })
  it('should calculate 0 NWPM from -1 GWPM and 1% accuracy', () => {
    expect(calculateNWPM(-1, 0.01)).toBe(0)
  })
  it('should calculate 0 NWPM from -1 GWPM and -1% accuracy', () => {
    expect(calculateNWPM(-1, -0.01)).toBe(0)
  })
})
