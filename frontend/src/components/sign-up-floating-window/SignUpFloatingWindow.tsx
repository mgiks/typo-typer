import './SignUpFloatingWindow.css'

function SignUpFloatingWindow(
  { shouldBeShown }: { shouldBeShown: boolean },
) {
  interface FormElements extends HTMLFormControlsCollection {
    name: HTMLInputElement
    email: HTMLInputElement
    password: HTMLInputElement
  }
  interface SignUpFormElement extends HTMLFormElement {
    readonly elements: FormElements
  }

  function submitForm(event: React.FormEvent<SignUpFormElement>) {
    event.preventDefault()

    const name = event.currentTarget.elements.name.value
    const email = event.currentTarget.elements.email.value
    const password = event.currentTarget.elements.password.value

    const data = { name, email, password }

    fetch('http://localhost:8000/users', {
      method: 'POST',
      body: JSON.stringify(data),
    }).then((responce) => console.log(responce))
  }

  const form = (
    <div id='form-wrapper'>
      Sign-up form:
      <form id='form' onSubmit={submitForm}>
        <label htmlFor='name'>Name:</label>
        <input type='text' id='name' />

        <label htmlFor='email'>Email:</label>
        <input type='email' id='email' />

        <label htmlFor='password'>Password:</label>
        <input type='password' id='password' />

        <button type='submit'>Sign up</button>
      </form>
    </div>
  )

  return shouldBeShown ? form : null
}

export default SignUpFloatingWindow
