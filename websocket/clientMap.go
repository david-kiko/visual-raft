package wsCore

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map //  key 是客户端IP  value 就是 WsClient连接对象
}

func (t *ClientMapStruct) Store(conn *websocket.Conn) {
	wsClient := NewWsClient(conn)
	t.data.Store(conn.RemoteAddr().String(), wsClient)
	//go wsClient.Ping(time.Second * 30)
	go wsClient.ReadLoop() //处理读 循环
}

func (t *ClientMapStruct) SendAll(v interface{}) {

	t.data.Range(func(key, value interface{}) bool {
		c := value.(*WsClient)
		c.Mux.Lock()
		defer c.Mux.Unlock()
		err := c.conn.WriteJSON(v)
		if err != nil {
			t.Remove(c.conn)
			log.Println(err)
		}
		return true
	})
}
func (t *ClientMapStruct) Remove(conn *websocket.Conn) {
	t.data.Delete(conn.RemoteAddr().String())
}
