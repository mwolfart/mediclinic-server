package handlers

import (
	"context"
	"fmt"
	"medclin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserPayload struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserHandler(c *gin.Context) {
	ctx := context.Background()

	users, err := models.Genericusers().All(ctx, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unknown error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func AddUserHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	var payload *CreateUserPayload
	err = c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload format is invalid",
		})
		return
	}

	user := &models.Genericuser{}
	user.Name = payload.Name
	user.Age = null.Int{Int: payload.Age, Valid: true}
	user.Gender = null.String{String: payload.Gender, Valid: true}

	// TODO validate e-mail
	user.Email = payload.Email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating hash from password",
		})
		return
	}
	user.Passwordhash = string(hashedPassword)

	err = user.Insert(c, db, boil.Infer())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
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
