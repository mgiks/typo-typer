import './ReminderToFocus.scss'

type ReminderToFocusProps = {
  focused: boolean
  focusSetter: (i: boolean) => void
  refToFocus: {
    current: {
      focus: () => void
    } | null
  }
}

function ReminderToFocus(
  { focused, focusSetter, refToFocus }: ReminderToFocusProps,
) {
  const reminderToFocus = (
    <div
      className='reminder-to-focus'
      data-testid='reminder-to-focus'
      onClick={() => (focusSetter(true), refToFocus.current?.focus())}
    >
      Click here or press any key to focus
    </div>
  )

  return (
    focused ? null : reminderToFocus
  )
}

export default ReminderToFocus
