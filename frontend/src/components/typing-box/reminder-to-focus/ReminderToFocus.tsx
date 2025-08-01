import './ReminderToFocus.scss'

type ReminderToFocusProps = {
  focused: boolean
  focusSetter: (i: boolean) => void
}

function ReminderToFocus({ focused, focusSetter }: ReminderToFocusProps) {
  const reminderToFocus = (
    <div
      className='reminder-to-focus'
      data-testid='reminder-to-focus'
      onClick={() => focusSetter(true)}
    >
      Click here or press any key to focus
    </div>
  )

  return (
    focused ? null : reminderToFocus
  )
}

export default ReminderToFocus
