import { calculateTypingAccuracy } from '../calculateTypingAccuracy'

describe('calculateTypingAccuracy', () => {
  it('should calculate 100% typing accuracy from 100 correct keypresses and 0 errors', () => {
    expect(calculateTypingAccuracy(100, 0)).toBe(1)
  })
  it('should calculate 50% typing accuracy from 50 correct keypresses and 50 errors', () => {
    expect(calculateTypingAccuracy(50, 50)).toBe(0.5)
  })
  it('should calculate 0% typing accuracy from 0 correct keypresses and 100 errors', () => {
    expect(calculateTypingAccuracy(0, 100)).toBe(0)
  })
  it('should calculate 0% typing accuracy from -1 correct keypresses and 100 errors', () => {
    expect(calculateTypingAccuracy(-1, 100)).toBe(0)
  })
  it('should calculate 0% typing accuracy from 100 correct keypresses and -1 errors', () => {
    expect(calculateTypingAccuracy(100, -1)).toBe(0)
  })
  it('should calculate 0% typing accuracy from -1 correct keypresses and -1 errors', () => {
    expect(calculateTypingAccuracy(-1, -1)).toBe(0)
  })
})
