package setup

import (
	"medclin/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var routes []*Route = []*Route{
	{Method: "POST", Path: "/login", Handler: handlers.LoginHandler},
	{Method: "GET", Path: "/users", Handler: handlers.GetUsersHandler},
	{Method: "POST", Path: "/users", Handler: handlers.AddUserHandler},
	{Method: "PATCH", Path: "/users/:id", Handler: handlers.UpdateUserHandler},
	{Method: "DELETE", Path: "/users/:id", Handler: handlers.DeleteUserHandler},
	{Method: "GET", Path: "/patient/insurances", Handler: handlers.GetOwnInsurances},
	{Method: "GET", Path: "/patient/:id/insurances", Handler: handlers.GetPatientInsurances},
	{Method: "POST", Path: "/patient/insurances", Handler: handlers.AddInsuranceToOwn},
	{Method: "POST", Path: "/patient/:id/insurances", Handler: handlers.AddInsuranceToPatient},
	{Method: "DELETE", Path: "/patient/insurances/:inId", Handler: handlers.RemoveOwnInsurance},
	{Method: "DELETE", Path: "/patient/:id/insurances/:inId", Handler: handlers.RemoveInsuranceFromPatient},
	{Method: "GET", Path: "/doctor/insurances", Handler: handlers.GetOwnSupportedInsurances},
	{Method: "GET", Path: "/doctor/:id/insurances", Handler: handlers.GetDoctorSupportedInsurances},
	{Method: "POST", Path: "/doctor/insurances", Handler: handlers.AddSupportedInsuranceToOwn},
	{Method: "POST", Path: "/doctor/:id/insurances", Handler: handlers.AddSupportedInsuranceToDoctor},
	{Method: "DELETE", Path: "/doctor/insurances/:inId", Handler: handlers.RemoveOwnSupportedInsurance},
	{Method: "DELETE", Path: "/doctor/:id/insurances/:inId", Handler: handlers.RemoveSupportedInsuranceFromDoctor},
	{Method: "GET", Path: "/doctor/specialties", Handler: handlers.GetOwnSpecialties},
	{Method: "GET", Path: "/doctor/:id/specialties", Handler: handlers.GetDoctorSpecialties},
	{Method: "POST", Path: "/doctor/specialties", Handler: handlers.AddSpecialtyToOwn},
	{Method: "POST", Path: "/doctor/:id/specialties", Handler: handlers.AddSpecialtyToDoctor},
	{Method: "DELETE", Path: "/doctor/specialties/:inId", Handler: handlers.RemoveOwnSpecialty},
	{Method: "DELETE", Path: "/doctor/:id/specialties/:inId", Handler: handlers.RemoveSpecialtyFromDoctor},
	{Method: "GET", Path: "/insurances", Handler: handlers.GetInsuranceHandler},
	{Method: "POST", Path: "/insurances", Handler: handlers.AddInsuranceHandler},
	{Method: "DELETE", Path: "/insurances/:id", Handler: handlers.DeleteInsuranceHandler},
	{Method: "GET", Path: "/roles", Handler: handlers.GetRoleHandler},
	{Method: "POST", Path: "/roles", Handler: handlers.AddRoleHandler},
	{Method: "DELETE", Path: "/roles/:id", Handler: handlers.DeleteRoleHandler},
}
