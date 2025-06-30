import { useIsSearchingForMatch } from '../../stores/MultiplayerStore'
import { PlayerModes, usePlayerMode } from '../../stores/TypingStatsStore'
import './GameSearchReminderCurtain.css'

function GameSearchReminderCurtain() {
  const playerMode = usePlayerMode()
  const isSearchingForPlayers = useIsSearchingForMatch()
  const gameSearchReminderCurtain = (
    <div id='game-search-reminder-curtain'>
      Press <kbd>Tab</kbd> to search for a game
    </div>
  )

  return playerMode === PlayerModes.Multiplayer && !isSearchingForPlayers
    ? gameSearchReminderCurtain
    : null
}

export default GameSearchReminderCurtain
