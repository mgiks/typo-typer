import { useRef } from 'react'
import './LogInFormFloatingWindow.css'

function LogInFormFloatingWindow(
  { shouldLogInFormBeShown }: { shouldLogInFormBeShown: boolean },
) {
  const formRef = useRef<HTMLFormElement>(null)

  function submitForm() {
    if (!formRef.current) return

    const data = new FormData(formRef.current)

    fetch('http://localhost:8000/users', {
      method: 'post',
      body: data,
    })
  }

  const logInForm = (
    <div id='log-in-form-wrapper'>
      Log in:
      <form ref={formRef} id='log-in-form' onSubmit={submitForm}>
        <label htmlFor='name'>Name:</label>
        <input type='text' id='name' />

        <label htmlFor='email'>Email:</label>
        <input type='email' id='email' />

        <label htmlFor='password'>Password:</label>
        <input type='password' id='password' />

        <input type='submit' value={'Log In'} />
      </form>
    </div>
  )

  return shouldLogInFormBeShown ? logInForm : null
}

export default LogInFormFloatingWindow
