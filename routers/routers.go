package routers

import (
	"keyValueProject/core/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		statusCode := c.Writer.Status()
		log := make(map[string]interface{})
		log["TimeStamp"] = time.Now()
		log["Method"] = c.Request.Method
		log["StatusCode"] = statusCode
		log["ClienIp"] = c.ClientIP()
		log["Path"] = c.FullPath()
		log["ErrorMessage"] = c.Errors
		if statusCode >= 200 && statusCode <= 204 {
			logging.Info(log)
		} else if statusCode >= 500 && statusCode <= 503 {
			logging.Error(log)
		} else {
			logging.Warning(log)
		}
	})
	router.GET("/keyValues", getKeyValues)
	router.POST("/keyValues", createKeyValue)
	router.GET("/keyValues/:key", getValueByKey)

	return router
}
