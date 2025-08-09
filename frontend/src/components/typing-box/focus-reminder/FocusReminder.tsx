import './FocusReminder.scss'

type FocusReminderProps = {
  visible: boolean
}

function FocusReminder({ visible }: FocusReminderProps) {
  const reminderToFocus = (
    <div className='focus-reminder' data-testid='focus-reminder'>
      Click here or press any key to focus
    </div>
  )

  return visible ? reminderToFocus : null
}

export default FocusReminder
