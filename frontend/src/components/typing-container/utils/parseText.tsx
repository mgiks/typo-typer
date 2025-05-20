export function parseText(text: string) {
  return text.replaceAll(/\r?\n|\r|\s+/g, ' ')
}
