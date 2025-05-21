export function keepCursorInView(
  cursor: HTMLSpanElement | null,
  textArea: HTMLDivElement | null,
) {
  if (!cursor || !textArea) {
    return
  }
  const cursorYPosition = cursor.getBoundingClientRect().y
  const textAreaTopYCoor = textArea.getBoundingClientRect().y
  const textAreaHeight = textArea.getBoundingClientRect().height
  const textAreaBottomYCoor = textAreaTopYCoor + textAreaHeight
  const isCursorInsideTextArea = textAreaTopYCoor < cursorYPosition &&
    cursorYPosition < textAreaBottomYCoor

  if (!isCursorInsideTextArea) {
    cursor.scrollIntoView({ behavior: 'smooth' })
  }
}
