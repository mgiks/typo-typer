import './ReminderToFocus.scss'

type ReminderToFocusProps = {
  visible: boolean
}

function ReminderToFocus({ visible }: ReminderToFocusProps) {
  const reminderToFocus = (
    <div className='reminder-to-focus' data-testid='reminder-to-focus'>
      Click here or press any key to focus
    </div>
  )

  return visible ? reminderToFocus : null
}

export default ReminderToFocus
