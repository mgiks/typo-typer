import './FocusReminder.scss'

type FocusReminderProps = {
  visible: boolean
}

function FocusReminder({ visible }: FocusReminderProps) {
  const reminderToFocus = (
    <div
      className='focus-reminder'
      role='status'
      aria-live='polite'
    >
      Click here or press any key to focus
    </div>
  )

  return visible ? reminderToFocus : null
}

export default FocusReminder
