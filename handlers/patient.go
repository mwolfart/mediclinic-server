package handlers

import (
	"medclin/models"
	"medclin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AddInsurancePayload struct {
	InsuranceId int `json:"insurance_id"`
}

func addInsuranceToPatientById(c *gin.Context, id int) {
	patient, err := models.Patients(qm.Where("userId = ?", id), qm.Load(models.PatientRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No patient data for the specified user")
		return
	}

	var payload *AddInsurancePayload
	err = c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid payload")
		return
	}

	insurance, err := models.Insurances(qm.Where("userId = ?", payload.InsuranceId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Insurance id specified could not be found")
		return
	}

	err = patient.AddInsuranceidInsurances(c, db, true, insurance)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusInternalServerError, "Error adding insurance")
		return
	}

	utils.SendData(c, patient)
}

func removeInsuranceFromPatientById(c *gin.Context, id int) {
	patient, err := models.Patients(qm.Where("userId = ?", id), qm.Load(models.PatientRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No patient data for the specified user")
		return
	}

	insuranceId := c.Param(("inId"))
	insurance, err := models.Insurances(qm.Where("userId = ?", insuranceId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Insurance id specified could not be found")
		return
	}

	err = patient.RemoveInsuranceidInsurances(c, db, insurance)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Error removing insurance from patient")
		return
	}

	utils.SendData(c, patient)
}

func GetOwnInsurances(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	patient, err := models.Patients(qm.Where("userId = ?", userId), qm.Load(models.PatientRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No patient data for the specified user")
		return
	}

	insurances := patient.R.InsuranceidInsurances
	utils.SendData(c, insurances)
}

func GetPatientInsurances(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	patientId := c.Param(("id"))
	patient, err := models.Patients(qm.Where("userId = ?", patientId), qm.Load(models.PatientRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No patient data for the specified user")
		return
	}

	insurances := patient.R.InsuranceidInsurances
	utils.SendData(c, insurances)
}

func AddInsuranceToOwn(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	addInsuranceToPatientById(c, userId)
}

func AddInsuranceToPatient(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	patientId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid patient ID")
		return
	}
	addInsuranceToPatientById(c, int(patientId))
}

func RemoveOwnInsurance(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	removeInsuranceFromPatientById(c, userId)
}

func RemoveInsuranceFromPatient(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	patientId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid patient ID")
		return
	}
	removeInsuranceFromPatientById(c, int(patientId))
}
