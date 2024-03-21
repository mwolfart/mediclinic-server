package handlers

import (
	"medclin/models"
	"medclin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AddSupportedInsurancePayload struct {
	InsuranceId int `json:"insurance_id"`
}

type AddSpecialtyPayload struct {
	SpecialtyId int `json:"specialty_id"`
}

func addInsuranceToDoctorById(c *gin.Context, id int) {
	doctor, err := models.Doctors(qm.Where("id = ?", id), qm.Load(models.DoctorRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	var payload *AddSupportedInsurancePayload
	err = c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid payload")
		return
	}

	insurance, err := models.Insurances(qm.Where("id = ?", payload.InsuranceId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Insurance id specified could not be found")
		return
	}

	err = doctor.AddInsuranceidInsurances(c, db, true, insurance)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusInternalServerError, "Error adding insurance")
		return
	}

	utils.SendData(c, doctor)
}

func removeInsuranceFromDoctorById(c *gin.Context, id int) {
	doctor, err := models.Doctors(qm.Where("id = ?", id), qm.Load(models.DoctorRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	insuranceId := c.Param(("inId"))
	insurance, err := models.Insurances(qm.Where("id = ?", insuranceId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Insurance id specified could not be found")
		return
	}

	err = doctor.RemoveInsuranceidInsurances(c, db, insurance)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Error removing insurance from doctor")
		return
	}

	utils.SendData(c, doctor)
}

func addSpecialtyToDoctorById(c *gin.Context, id int) {
	doctor, err := models.Doctors(qm.Where("id = ?", id), qm.Load(models.DoctorRels.SpecialtyidSpecialties)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	var payload *AddSpecialtyPayload
	err = c.BindJSON(&payload)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid payload")
		return
	}

	specialty, err := models.Specialties(qm.Where("id = ?", payload.SpecialtyId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Specialty id specified could not be found")
		return
	}

	err = doctor.AddSpecialtyidSpecialties(c, db, true, specialty)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusInternalServerError, "Error adding specialty")
		return
	}

	utils.SendData(c, doctor)
}

func removeSpecialtyFromDoctorById(c *gin.Context, id int) {
	doctor, err := models.Doctors(qm.Where("id = ?", id), qm.Load(models.DoctorRels.SpecialtyidSpecialties)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	specialtyId := c.Param(("spId"))
	specialty, err := models.Specialties(qm.Where("id = ?", specialtyId)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Specialty id specified could not be found")
		return
	}

	err = doctor.RemoveSpecialtyidSpecialties(c, db, specialty)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "Error removing specialty from doctor")
		return
	}

	utils.SendData(c, doctor)
}

func GetOwnSupportedInsurances(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	doctor, err := models.Doctors(qm.Where("id = ?", userId), qm.Load(models.DoctorRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	insurances := doctor.R.InsuranceidInsurances
	utils.SendData(c, insurances)
}

func GetDoctorSupportedInsurances(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	doctorId := c.Param(("id"))
	doctor, err := models.Doctors(qm.Where("id = ?", doctorId), qm.Load(models.DoctorRels.InsuranceidInsurances)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	insurances := doctor.R.InsuranceidInsurances
	utils.SendData(c, insurances)
}

func AddSupportedInsuranceToOwn(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	addInsuranceToDoctorById(c, userId)
}

func AddSupportedInsuranceToDoctor(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	doctorId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}
	addInsuranceToDoctorById(c, int(doctorId))
}

func RemoveOwnSupportedInsurance(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	removeInsuranceFromDoctorById(c, userId)
}

func RemoveSupportedInsuranceFromDoctor(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	doctorId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}
	removeInsuranceFromDoctorById(c, int(doctorId))
}

func GetOwnSpecialties(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	doctor, err := models.Doctors(qm.Where("id = ?", userId), qm.Load(models.DoctorRels.SpecialtyidSpecialties)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	specialties := doctor.R.SpecialtyidSpecialties
	utils.SendData(c, specialties)
}

func GetDoctorSpecialties(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	doctorId := c.Param(("id"))
	doctor, err := models.Doctors(qm.Where("id = ?", doctorId), qm.Load(models.DoctorRels.SpecialtyidSpecialties)).One(c, db)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusNotFound, "No doctor data for the specified user")
		return
	}

	specialties := doctor.R.SpecialtyidSpecialties
	utils.SendData(c, specialties)
}

func AddSpecialtyToOwn(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	addSpecialtyToDoctorById(c, userId)
}

func AddSpecialtyToDoctor(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	doctorId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}
	addSpecialtyToDoctorById(c, int(doctorId))
}

func RemoveOwnSpecialty(c *gin.Context) {
	userId, err := utils.UserEmailFromToken(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, "Invalid authorization token")
		return
	}

	removeSpecialtyFromDoctorById(c, userId)
}

func RemoveSpecialtyFromDoctor(c *gin.Context) {
	ok, err := HasPermission(c, "manager")
	if !ok || err != nil {
		utils.SendErrorMessage(c, http.StatusForbidden, "Not authorized to execute this operation")
		return
	}

	idParam := c.Param(("id"))
	doctorId, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusBadRequest, "Invalid doctor ID")
		return
	}
	removeSpecialtyFromDoctorById(c, int(doctorId))
}
