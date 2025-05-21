export function calculateWordCount(text: string) {
  text = text.trim()
  text = text.replaceAll(/\s+/g, ' ')
  if (!text) return 0
  return text.split(' ').length
}
