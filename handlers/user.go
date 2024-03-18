package handlers

import (
	"fmt"
	"medclin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	users, err := models.Genericusers().AllG(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unknown error: %s", err),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
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
