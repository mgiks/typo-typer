import { MessageDataMap, MessageType } from './Message'

export type MessageOf<T extends MessageType> = {
  type: T
  data: MessageDataMap[T]
}

export function NewMessage<T extends MessageType>(
  messageType: T,
  data: MessageDataMap[T],
): MessageOf<T> {
  return {
    type: messageType,
    data,
  }
}
