package handlers

import (
	"fmt"
	"medclin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type CreateInsurancePayload struct {
	Name string `json:"name"`
}

func GetInsuranceHandler(c *gin.Context) {
	insurances, err := models.Insurances().All(c, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Unknown error: %s", err),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": insurances,
	})
}

func AddInsuranceHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	var payload *CreateInsurancePayload
	err = c.BindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Payload format is invalid",
		})
		return
	}

	insurance := &models.Insurance{}
	insurance.Name = payload.Name

	err = insurance.Insert(c, db, boil.Infer())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": insurance,
	})
}

func DeleteInsuranceHandler(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Not authorized to execute this operation",
		})
		return
	}

	id := c.Param("id")
	insurance, err := models.Insurances(qm.Where("id = ?", id)).One(c, db)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found",
		})
		return
	}

	_, err = insurance.Delete(c, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't delete the specified insurance. It may belong to one or more users.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
