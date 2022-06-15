package wsCore

type WsMessage struct {
	MessageType int
	MessageData []byte //  一般是JSON格式
}

func NewWsMessage(messageType int, messageData []byte) *WsMessage {
	return &WsMessage{MessageType: messageType, MessageData: messageData}
}
