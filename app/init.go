package app

import (
	"net/http"
	"xiaochat/lib/logs"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer() {
	web := gin.Default()
	web.Use(logs.Logger())

	web.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "LspChat server!")
	})
	web.GET("/ping", ping)

	if err := web.Run(":6969"); err != nil {
		logs.DefaultLog("startup server failed, err: %v\n", err)
	}
}

func ping(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
