package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInsuranceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}

func AddInsuranceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}

func DeleteInsuranceHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of insurances",
	})
}
