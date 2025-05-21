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
import { Message, MessageType } from './types/Message'
import { parseText } from './utils/parseText'
import { MessageOf, NewMessage } from './types/MessageInitializers'

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

  const playerId = usePlayerId()
  const playerName = usePlayerName()
  useEffect(() => {
    if (!isSearchingForMatch) return

    const ws = new WebSocket('ws://localhost:8000')

    ws.onopen = () => {
      console.log('Connected to websocket server')

      const message = NewMessage(MessageType.searchForMatch, {
        playerName,
        playerId,
      })

      ws.send(JSON.stringify(message))
    }

    ws.onmessage = (event) => {
      const message: Message = JSON.parse(event.data)

      switch (message.type) {
        case MessageType.matchFound:
          setMatchFoundData(message)
          break
      }
    }

    ws.onclose = () => {
      stopSearchingForMatch()
      console.log('Closed websocket connection')
    }

    return () => ws.close()
  }, [isSearchingForMatch])

  const lastPressedKey = useRef('')
  useEffect(() => {
    wrongTextStartIndex > -1 && lastPressedKey.current !== 'Backspace'
      ? increaseWrongKeyCount()
      : increaseCorrectKeyCount()
  }, [wrongTextStartIndex, cursorIndex])

  useEffect(() => {
    const textBeforeCursor = text.split('').slice(0, cursorIndex).join('')
    const textAfterCursor = text.split('').slice(cursorIndex).join('')

    setTextAfterCursor(textAfterCursor)
    setTextBeforeCursor(textBeforeCursor)

    const isAtTheEndOfText = textBeforeCursor && !textAfterCursor
    isAtTheEndOfText && finishTypingGame()
  })

  async function getRandomText() {
    const response = await fetch('http://localhost:8000/random-texts')
    const randomTextMessage = await response.json() as MessageOf<
      MessageType.randomText
    >
    return randomTextMessage.data.text
  }
  useEffect(() => {
    getRandomText().then((text) => setText(parseText(text)))
  }, [textRefreshCount])

  useEffect(() => {
    cursorIndex > 0 && setCursorToMoved()
  }, [cursorIndex])

  useEffect(() => {
    cursorIndex <= wrongTextStartIndex && setWrongTextStartIndex(-1)
  }, [cursorIndex, wrongTextStartIndex])

  function updateCursorIndexByKey(key: string) {
    key === 'Backspace' ? decreaseCursorIndex() : increaseCursorIndex()
  }

  function checkIfKeyIsWrong(key: string) {
    return !(text.at(cursorIndex) === key)
  }

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
