import { usePlayerName } from '../../stores/MultiplayerStore'
import './HeaderContainer.css'

function HeaderContainer(
  { setShouldLogInFormBeShown }: {
    setShouldLogInFormBeShown: React.Dispatch<React.SetStateAction<boolean>>
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
          id='log-in-button'
          onClick={() => {
            setShouldLogInFormBeShown((prev) => !prev)
          }}
        >
          Log in
        </button>
      </div>
    </div>
  )
}

export default HeaderContainer
