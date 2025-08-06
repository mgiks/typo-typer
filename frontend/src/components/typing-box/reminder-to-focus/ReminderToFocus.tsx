import './ReminderToFocus.scss'

type ReminderToFocusProps = {
  focused: boolean
}

function ReminderToFocus({ focused }: ReminderToFocusProps) {
  const reminderToFocus = (
    <div
      className='reminder-to-focus'
      data-testid='reminder-to-focus'
    >
      Click here or press any key to focus
    </div>
  )

  return (
    focused ? null : reminderToFocus
  )
}

export default ReminderToFocus
