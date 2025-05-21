export function extractCorrectText(
  textBeforeCursor: string,
  wrongTextStartIndex: number,
) {
  return wrongTextStartIndex > -1
    ? textBeforeCursor.slice(0, wrongTextStartIndex)
    : textBeforeCursor
}
