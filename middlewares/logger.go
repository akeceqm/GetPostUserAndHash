package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() {
	server := gin.New()

	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	server.Use(gin.Recovery())

}

func LogFile(c *gin.Engine) {
	file, err := os.Create("logger.log")
	if err != nil {
		log.Println("Error creating log file:", err)
		return
	}
	c.Use(gin.LoggerWithWriter(file))
}
