export function extractWrongText(
  textBeforeCursor: string,
  wrongTextStartIndex: number,
) {
  return wrongTextStartIndex > -1
    ? textBeforeCursor.slice(wrongTextStartIndex)
    : ''
}
