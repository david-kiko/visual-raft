package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
	"net/http"
	"os"
	wsCore "visual-raft/websocket"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(CorsMiddleware())
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.GET("/ws", wsCore.ServeWs)
	r.POST("/send", send)
	httpAddr := ":8000"
	log.Info().Msgf("web界面监听于本机端口 %s", httpAddr)

	err := r.Run(httpAddr)
	if err != nil {
		log.Info().Msgf("本地端口已被占用,可能已有web实例启动 %s", err)
	}
}

func send(c *gin.Context) {
	body := make(map[string]interface{})
	c.ShouldBindJSON(&body)
	wsCore.ClientMap.SendAll(body)
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//直接取前端的地址允许跨域
		allow := c.Request.Header.Get("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", allow)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Request.Header.Del("Origin")
		c.Next()
	}
}

func init() {
	if !log.IsTerminal(os.Stderr.Fd()) {
		return
	}
	log.DefaultLogger = log.Logger{
		TimeFormat: "15:04:05",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		},
	}
}
