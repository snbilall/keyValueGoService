package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/keyValues", getKeyValues)
	router.POST("/keyValues", createKeyValue)
	router.GET("/keyValues/:key", getValueByKey)

	return router
}
