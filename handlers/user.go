package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
	})
}

func AddUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
	})
}

func UpdateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
	})
}

func DeleteUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
	})
}
