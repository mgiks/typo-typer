import { useEffect, useRef, useState } from 'react'
import './FocusReminder.scss'

export const FOCUS_REMINDER_TIMEOUT_MS = 750

type FocusReminderProps = {
  show: boolean
}

function FocusReminder({ show }: FocusReminderProps) {
  const [visible, setVisible] = useState(false)
  const timeOutRef = useRef(-1)

  useEffect(() => {
    if (show) {
      timeOutRef.current = window.setTimeout(
        setVisible,
        FOCUS_REMINDER_TIMEOUT_MS,
        true,
      )
    }

    return () => (setVisible(false), clearTimeout(timeOutRef.current))
  }, [show])

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
