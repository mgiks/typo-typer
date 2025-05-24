import { useRef } from 'react'
import './SignUpFloatingWindow.css'

function SignUpFloatingWindow(
  { shouldBeShown }: { shouldBeShown: boolean },
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

  const form = (
    <div id='form-wrapper'>
      Sign-up form:
      <form ref={formRef} id='form' onSubmit={submitForm}>
        <label htmlFor='name'>Name:</label>
        <input type='text' id='name' />

        <label htmlFor='email'>Email:</label>
        <input type='email' id='email' />

        <label htmlFor='password'>Password:</label>
        <input type='password' id='password' />

        <input type='submit' value={'Sign up'} />
      </form>
    </div>
  )

  return shouldBeShown ? form : null
}

export default SignUpFloatingWindow
