package wsCore

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/phuslu/log"
	"net/http"
)

var (
	UpGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func ServeWs(c *gin.Context) {
	ws, err := UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error().Msgf("ws握手失败 %s", err)
		return
	}
	log.Info().Msgf("新建ws连接来自 %s", ws.RemoteAddr())
	ClientMap.Store(ws)
}
