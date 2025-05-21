export function isControlKey(key: string) {
  return key !== 'Backspace' && key.length > 1
}
