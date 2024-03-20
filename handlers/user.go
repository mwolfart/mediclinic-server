package handlers

import (
	"context"
	"fmt"
	"medclin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserPayload struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserPayload struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
	Email  string `json:"email,omitempty"`
}

func GetUsersHandler(c *gin.Context) {
	ctx := context.Background()

	users, err := models.Genericusers(qm.Select("id", "name", "gender", "age")).All(ctx, db)
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
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
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

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUserHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	id := c.Param("id")
	var payload *UpdateUserPayload
	err = c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload format is invalid",
		})
		return
	}

	user, err := models.Genericusers(qm.Where("id = ?", id)).One(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	if payload.Email != "" {
		user.Email = payload.Email
	}
	if payload.Age != 0 {
		user.Age = null.Int{Int: payload.Age, Valid: true}
	}
	if payload.Gender != "" {
		user.Gender = null.String{String: payload.Gender, Valid: true}
	}
	if payload.Name != "" {
		user.Name = payload.Name
	}
	user.Update(c, db, boil.Infer())

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func DeleteUserHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	id := c.Param("id")
	user, err := models.Genericusers(qm.Where("id = ?", id)).One(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	user.Delete(c, db)
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
