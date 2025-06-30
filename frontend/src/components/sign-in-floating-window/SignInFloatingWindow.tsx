import './SignInFloatingWindow.css'

function SignInFloatingWindow(
  { shouldBeShown }: { shouldBeShown: boolean },
) {
  interface FormElements extends HTMLFormControlsCollection {
    name: HTMLInputElement
    email: HTMLInputElement
    password: HTMLInputElement
  }
  interface SignInFormElement extends HTMLFormElement {
    readonly elements: FormElements
  }

  function submitForm(event: React.FormEvent<SignInFormElement>) {
    event.preventDefault()

    const name = event.currentTarget.elements.name.value
    const email = event.currentTarget.elements.email.value
    const password = event.currentTarget.elements.password.value

    const data = { name, email, password }

    type jwtData = {
      accessToken: string
      tokenType: string
      expiresIn: number
    }

    fetch('http://localhost:8000/auth/signin', {
      method: 'POST',
      body: JSON.stringify(data),
    }).then((res) => res.json()).then((res) => res as jwtData).then((data) =>
      console.log(data)
    )
  }

  const form = (
    <div id='form-wrapper'>
      Sign-in form:
      <form id='form' onSubmit={submitForm}>
        <label htmlFor='name'>Name:</label>
        <input type='text' id='name' />

        <label htmlFor='email'>Email:</label>
        <input type='email' id='email' />

        <label htmlFor='password'>Password:</label>
        <input type='password' id='password' />

        <button type='submit'>Sign in</button>
      </form>
    </div>
  )

  return shouldBeShown ? form : null
}

export default SignInFloatingWindow
