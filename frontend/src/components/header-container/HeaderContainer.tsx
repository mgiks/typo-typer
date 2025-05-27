import { usePlayerName } from '../../stores/MultiplayerStore'
import './HeaderContainer.css'

function HeaderContainer(
  { shouldSignUpFormBeShownSetter, shouldSignInFormBeShownSetter }: {
    shouldSignUpFormBeShownSetter: React.Dispatch<
      React.SetStateAction<boolean>
    >
    shouldSignInFormBeShownSetter: React.Dispatch<
      React.SetStateAction<boolean>
    >
  },
) {
  const playerName = usePlayerName()

  return (
    <div id='header-container'>
      <div>
        {playerName}
      </div>
      <div>
        <button
          onClick={() => {
            shouldSignUpFormBeShownSetter((prev) => !prev)
          }}
        >
          Sign Up
        </button>
        <button
          onClick={() => {
            shouldSignInFormBeShownSetter((prev) => !prev)
          }}
        >
          Sign In
        </button>
      </div>
    </div>
  )
}

export default HeaderContainer
