package wsCore

import (
	"github.com/gorilla/websocket"
	"time"
)

type WsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage //读队列 (chan)
	closeChan chan byte       // 失败队列
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn, readChan: make(chan *WsMessage), closeChan: make(chan byte)}
}
func (t *WsClient) Ping(wait time.Duration) {
	for {
		time.Sleep(wait)
		err := t.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			ClientMap.Remove(t.conn)
			return
		}
	}
}
func (t *WsClient) ReadLoop() {
	for {
		messageType, data, err := t.conn.ReadMessage()
		if err != nil {
			t.conn.Close()
			ClientMap.Remove(t.conn)
			t.closeChan <- 1
			break
		}
		t.readChan <- NewWsMessage(messageType, data)
	}
}
