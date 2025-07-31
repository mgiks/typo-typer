type ReminderToFocusProps = {
  focused: boolean
}

function ReminderToFocus({ focused }: ReminderToFocusProps) {
  if (focused) {
    return null
  }

  return <div data-testid='reminder-to-focus'></div>
}

export default ReminderToFocus
