package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"message": message,
	})
}

func SendData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
