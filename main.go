package main

import (
	"keyValueProject/models"
	"keyValueProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	service.InitKeyValueObject()
	router := gin.Default()
	router.GET("/keyValues", getKeyValues)
	router.GET("/createKeyValue", createKeyValue)

	router.Run("localhost:8080")
}

func getKeyValues(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.KeyValueObject)
}

func createKeyValue(c *gin.Context) {
	var newKeyValue models.KeyValueRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newKeyValue); err != nil {
		return
	}

	service.AddKeyValue(newKeyValue.Key, newKeyValue.Value)
	c.IndentedJSON(http.StatusCreated, newKeyValue)
}
