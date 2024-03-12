package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}

func AddRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}

func DeleteRoleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}
