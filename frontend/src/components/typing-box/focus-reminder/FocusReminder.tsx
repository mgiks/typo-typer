import { useEffect, useRef, useState } from 'react'
import './FocusReminder.scss'

export const FOCUS_REMINDER = {
  TIMEOUT_MS: 750,
  TEXT: 'Click here or press any key to focus',
}

type FocusReminderProps = {
  show: boolean
}

function FocusReminder({ show }: FocusReminderProps) {
  const [visible, setVisible] = useState(false)
  const timeOutRef = useRef(-1)

  useEffect(() => {
    if (show) {
      timeOutRef.current = window.setTimeout(
        () => setVisible(true),
        FOCUS_REMINDER.TIMEOUT_MS,
      )
    }

    return () => (clearTimeout(timeOutRef.current), setVisible(false))
  }, [show])

  const reminderToFocus = (
    <div
      className='focus-reminder'
      role='status'
      aria-live='polite'
    >
      {FOCUS_REMINDER.TEXT}
    </div>
  )

  return visible ? reminderToFocus : null
}

export default FocusReminder
