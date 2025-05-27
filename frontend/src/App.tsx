import ResultContainer from './components/result-container/ResultContainer'
import TypingContainer from './components/typing-container/TypingContainer'
import TypingStatsContainer from './components/typing-stats-container/TypingStatsContainer'
import './App.css'
import {
  PlayerModes,
  useIsDoneTyping,
  usePlayerMode,
  useTypingStatsActions,
} from './stores/TypingStatsStore'
import { useEffect, useState } from 'react'
import { useTextActions } from './stores/TextStore'
import { useMultiplayerActions } from './stores/MultiplayerStore'
import HeaderContainer from './components/header-container/HeaderContainer'
import MatchProgressContainer from './components/match-progress-container/MatchProgressContainer'
import SignUpFloatingWindow from './components/sign-up-floating-window/SignUpFloatingWindow'
import SignInFloatingWindow from './components/sign-in-floating-window/SignInFloatingWindow'

function App() {
  const [shouldSignUpFormBeShown, setShouldSignUpFormBeShown] = useState(false)
  const [shouldSignInFormBeShown, setShouldSignInFormBeShown] = useState(false)

  const isDoneTyping = useIsDoneTyping()
  const playerMode = usePlayerMode()

  const { searchForMatch, stopSearchingForMatch } = useMultiplayerActions()
  const { increaseTextRefreshCount, resetCursorIndex } = useTextActions()
  const { resetTypingStats } = useTypingStatsActions()
  const { setRandomName, setRandomPlayerId } = useMultiplayerActions()

  function handleTab(event: KeyboardEvent) {
    if (event.key !== 'Tab') {
      return
    }
    switch (playerMode) {
      case PlayerModes.Singleplayer:
        increaseTextRefreshCount()

        resetTypingStats()
        resetCursorIndex()

        stopSearchingForMatch()
        break
      case PlayerModes.Multiplayer:
        searchForMatch()
        break
    }
  }
  useEffect(() => {
    const removeTabDefaultFunctionality = (event: KeyboardEvent) => {
      event.key === 'Tab' && event.preventDefault()
    }
    window.onkeydown = removeTabDefaultFunctionality

    window.onkeyup = handleTab
  }, [playerMode])

  useEffect(() => {
    setRandomName()
    setRandomPlayerId()
  }, [])

  return (
    <>
      <HeaderContainer
        shouldSignUpFormBeShownSetter={setShouldSignUpFormBeShown}
        shouldSignInFormBeShownSetter={setShouldSignInFormBeShown}
      />
      <SignInFloatingWindow shouldBeShown={shouldSignInFormBeShown} />
      <SignUpFloatingWindow shouldBeShown={shouldSignUpFormBeShown} />
      <MatchProgressContainer />
      <div
        id='typing-container-with-stats'
        className={isDoneTyping ? 'invisible' : undefined}
      >
        <TypingContainer isAnyFormShown={shouldSignUpFormBeShown} />
        <TypingStatsContainer />
      </div>
      <ResultContainer />
    </>
  )
}

export default App
