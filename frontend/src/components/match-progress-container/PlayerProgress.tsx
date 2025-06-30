import './PlayerProgress.css'

function PlayerProgress({ playerName }: { playerName: string }) {
  return (
    <div className='player-progress'>
      <div>{playerName}</div>
    </div>
  )
}

export default PlayerProgress
