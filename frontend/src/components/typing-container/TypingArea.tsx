import { KeyboardEvent, useEffect, useRef } from 'react'
import './TypingArea.css'
import { isControlKey } from './utils/isControlKey'
import {
  useCursorIndex,
  useText,
  useTextActions,
  useTextRefreshCount,
  useWrongTextStartIndex,
} from '../../stores/TextStore'
import { useTypingStatsActions } from '../../stores/TypingStatsStore'
import {
  useIsSearchingForMatch,
  useMultiplayerActions,
  usePlayerId,
  usePlayerName,
} from '../../stores/MultiplayerStore'
import { getRandomText } from './utils/getRandomText'
import {
  MatchFoundMessage,
  Message,
  SearchForMatchMessage,
} from './dtos/Message'
import { parseText } from './utils/parseText'

function TypingArea(
  { ref }: { ref: React.RefObject<HTMLTextAreaElement | null> },
) {
  const textRefreshCount = useTextRefreshCount()
  const cursorIndex = useCursorIndex()
  const isSearchingForMatch = useIsSearchingForMatch()
  const text = useText()
  const wrongTextStartIndex = useWrongTextStartIndex()
  const {
    setTextBeforeCursor,
    setTextAfterCursor,
    setWrongTextStartIndex,
    setText,
    increaseCursorIndex,
    decreaseCursorIndex,
  } = useTextActions()
  const {
    finishTypingGame,
    increaseWrongKeyCount,
    increaseCorrectKeyCount,
    setCursorToMoved,
  } = useTypingStatsActions()
  const {
    stopSearchingForMatch,
    setMatchFoundData,
  } = useMultiplayerActions()

  useEffect(() => {
    getRandomText().then((text) => setText(parseText(text)))
  }, [textRefreshCount])

  const playerId = usePlayerId()
  const playerName = usePlayerName()
  useEffect(() => {
    if (!isSearchingForMatch) return
    const ws = new WebSocket('ws://localhost:8000')

    ws.onopen = () => {
      console.log('Connected to websocket server')
      const playerInfo: SearchForMatchMessage = {
        type: 'searchForMatch',
        data: { playerName, playerId },
      }
      ws.send(JSON.stringify(playerInfo))
    }

    ws.onmessage = (event) => {
      const data: Message = JSON.parse(event.data)
      switch (data.type) {
        case 'matchFound':
          const matchFoundMessage = data as MatchFoundMessage
          setMatchFoundData(matchFoundMessage.data)
          break
        case 'gameUpdate':
          break
      }
    }

    ws.onclose = () => {
      console.log('Closed websocket connection')
      stopSearchingForMatch()
    }

    return () => ws.close()
  }, [isSearchingForMatch])

  useEffect(() => {
    cursorIndex > 0 && setCursorToMoved()
  }, [cursorIndex])

  useEffect(() => {
    const textBeforeCursor = text.split('').slice(0, cursorIndex).join('')
    const textAfterCursor = text.split('').slice(cursorIndex).join('')
    setTextAfterCursor(textAfterCursor)
    setTextBeforeCursor(textBeforeCursor)
    const isAtTheEndOfText = textBeforeCursor && !textAfterCursor
    isAtTheEndOfText && finishTypingGame()
  })

  function updateCursorIndexByKey(key: string) {
    key === 'Backspace' ? decreaseCursorIndex() : increaseCursorIndex()
  }

  function checkIfKeyIsWrong(key: string) {
    return !(text.at(cursorIndex) === key)
  }

  const lastPressedKey = useRef('')
  useEffect(() => {
    wrongTextStartIndex > -1 && lastPressedKey.current !== 'Backspace'
      ? increaseWrongKeyCount()
      : increaseCorrectKeyCount()
  }, [wrongTextStartIndex, cursorIndex])

  useEffect(() => {
    if (cursorIndex <= wrongTextStartIndex) setWrongTextStartIndex(-1)
  }, [cursorIndex, wrongTextStartIndex])

  function handleKeypress(event: KeyboardEvent) {
    const { key } = event

    if (isControlKey(key)) return
    lastPressedKey.current = key
    updateCursorIndexByKey(key)

    if (key === 'Backspace') return
    if (checkIfKeyIsWrong(key) && wrongTextStartIndex === -1) {
      setWrongTextStartIndex(cursorIndex)
    }
  }

  return (
    <textarea
      ref={ref}
      onKeyDown={handleKeypress}
      id='typing-area'
      autoFocus
    />
  )
}

export default TypingArea
