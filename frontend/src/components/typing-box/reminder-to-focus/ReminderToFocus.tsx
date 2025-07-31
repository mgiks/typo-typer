import './ReminderToFocus.scss'

type ReminderToFocusProps = {
  focused: boolean
}

function ReminderToFocus({ focused }: ReminderToFocusProps) {
  if (focused) {
    return null
  }

  return (
    <div className='reminder-to-focus' data-testid='reminder-to-focus'>
      Click here or press any key to focus
    </div>
  )
}

export default ReminderToFocus
