import { usePlayerName } from '../../stores/MultiplayerStore'
import './HeaderContainer.css'

function HeaderContainer(
  { shouldSignUpFromBeShownSetter }: {
    shouldSignUpFromBeShownSetter: React.Dispatch<
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
            shouldSignUpFromBeShownSetter((prev) => !prev)
          }}
        >
          Sign Up
        </button>
      </div>
    </div>
  )
}

export default HeaderContainer
