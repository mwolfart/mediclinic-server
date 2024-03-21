package handlers

import (
	"fmt"
	"medclin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateRolePayload struct {
	Description string `json:"description"`
}

func GetRoleHandler(c *gin.Context) {
	roles, err := models.Roles().All(c, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unknown error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})
}

func AddRoleHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	var payload *CreateRolePayload
	err = c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload format is invalid",
		})
		return
	}

	role := &models.Role{}
	role.Description = payload.Description

	err = role.Insert(c, db, boil.Infer())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}

func DeleteRoleHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	id := c.Param("id")
	role, err := models.Roles(qm.Where("id = ?", id)).One(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	_, err = role.Delete(c, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't delete the specified role. It may belong to one or more users.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
