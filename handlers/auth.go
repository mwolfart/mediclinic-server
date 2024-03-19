package handlers

import (
	"context"
	"medclin/models"
	"medclin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	ctx := context.Background()

	var payload *LoginPayload
	err := c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload format is invalid. Required: (email, password)",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
	}

	user, err := models.Genericusers(qm.Where("email = ? AND passwordHash = ?", payload.Email, string(hash))).One(ctx, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
		})
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error fetching token",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
