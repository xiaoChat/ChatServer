package logs

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	Default = "0"
	Info    = "1"
	Warn    = "2"
	Error   = "4"
)

func DefaultLog(format string, v ...interface{}) {
	// 后续落文件
	log.Printf(`[]`+format, v...)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
