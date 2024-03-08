package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gin() {
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List of users",
		})
	})
	r.Run()
}
