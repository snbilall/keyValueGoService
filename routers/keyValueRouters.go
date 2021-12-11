package routers

import (
	"keyValueProject/models"
	"keyValueProject/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func getKeyValues(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.KeyValueObject)
}

func createKeyValue(c *gin.Context) {
	var newKeyValue models.KeyValueRequest

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newKeyValue); err != nil {
		return
	}
	validate := validator.New()
	if err := validate.Struct(newKeyValue); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	service.AddKeyValue(newKeyValue)
	c.IndentedJSON(http.StatusCreated, newKeyValue)
}

func getValueByKey(c *gin.Context) {
	key := c.Param("key")

	value, err := service.GetKeyValue(key)
	if err == nil {
		c.IndentedJSON(http.StatusOK, value)
	} else {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	}
}
