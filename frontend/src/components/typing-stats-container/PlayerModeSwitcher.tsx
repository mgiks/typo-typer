import singleplayerIcon from './icons/singleplayer-icon.svg'
import multiplayerIcon from './icons/multiplayer-icon.svg'
import privateRoomIcon from './icons/private-room-icon.svg'
import './PlayerModeSwitcher.css'
import { useState } from 'react'
import {
  PlayerModes,
  usePlayerMode,
  useTypingStatsActions,
} from '../../stores/TypingStatsStore'

export const playerModeToIcon: Map<PlayerModes, string> = new Map([
  [PlayerModes.Singleplayer, singleplayerIcon],
  [PlayerModes.Multiplayer, multiplayerIcon],
  [PlayerModes.PrivateRoom, privateRoomIcon],
])

function PlayerModeSwitcher() {
  const playerMode = usePlayerMode()
  const [isSwitcherVisible, setIsSwitcherVisible] = useState(false)

  function toggleSwitcherVisibility() {
    setIsSwitcherVisible((visibilityValue) => !visibilityValue)
  }

  return (
    <div>
      <button
        onClick={() => toggleSwitcherVisibility()}
      >
        <img
          className='player-mode-icon'
          src={playerModeToIcon.get(playerMode)}
        />
      </button>
      <div
        id='player-mode-dropdown'
        className={isSwitcherVisible ? 'visible' : 'invisible'}
      >
        <ul>
          <PlayerMode
            playerMode={PlayerModes.Multiplayer}
            setIsSwitcherVisible={setIsSwitcherVisible}
          />
          <PlayerMode
            playerMode={PlayerModes.Singleplayer}
            setIsSwitcherVisible={setIsSwitcherVisible}
          />
          <PlayerMode
            playerMode={PlayerModes.PrivateRoom}
            setIsSwitcherVisible={setIsSwitcherVisible}
          />
        </ul>
      </div>
    </div>
  )
}

function PlayerMode(
  { playerMode, setIsSwitcherVisible }: {
    playerMode: PlayerModes
    setIsSwitcherVisible: React.Dispatch<React.SetStateAction<boolean>>
  },
) {
  const { setPlayerMode } = useTypingStatsActions()

  return (
    <li>
      <button
        onClick={() => (setPlayerMode(playerMode),
          setIsSwitcherVisible((visibilityValue) => !visibilityValue))}
      >
        <img
          className='player-mode-icon'
          src={playerModeToIcon.get(playerMode)}
        />{' '}
        {playerMode}
      </button>
    </li>
  )
}

export default PlayerModeSwitcher
