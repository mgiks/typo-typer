import './LogInFormFloatingWindow.css'

function LogInFormFloatingWindow() {
  return (
    <div id='log-in-form-wrapper'>
      Log in:
      <form id='log-in-form'>
        <label htmlFor='name'>Name:</label>
        <input type='text' id='name' />

        <label htmlFor='email'>Email:</label>
        <input type='email' id='email' />

        <label htmlFor='password'>Password:</label>
        <input type='password' id='password' />
      </form>
    </div>
  )
}

export default LogInFormFloatingWindow
