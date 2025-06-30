import { useMatchFoundData } from '../../stores/MultiplayerStore'
import './MatchProgressContainer.css'
import PlayerProgress from './PlayerProgress'

function MatchProgressContainer() {
  const matchFoundData = useMatchFoundData()

  if (!matchFoundData) return null

  const playersProgress = []
  for (const playerName of matchFoundData.playerNames) {
    playersProgress.push(
      <PlayerProgress playerName={playerName} />,
    )
  }

  return (
    <div id='match-progress-container'>
      {playersProgress}
    </div>
  )
}

export default MatchProgressContainer
