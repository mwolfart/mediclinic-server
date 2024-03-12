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
	{Method: "GET", Path: "/users", Handler: handlers.GetUserHandler},
	{Method: "POST", Path: "/users", Handler: handlers.AddUserHandler},
	{Method: "PATCH", Path: "/users", Handler: handlers.UpdateUserHandler},
	{Method: "DELETE", Path: "/users", Handler: handlers.DeleteUserHandler},
	{Method: "GET", Path: "/insurances", Handler: handlers.GetInsuranceHandler},
	{Method: "POST", Path: "/insurances", Handler: handlers.AddInsuranceHandler},
	{Method: "DELETE", Path: "/insurances", Handler: handlers.DeleteInsuranceHandler},
	{Method: "GET", Path: "/roles", Handler: handlers.GetRoleHandler},
	{Method: "POST", Path: "/roles", Handler: handlers.AddRoleHandler},
	{Method: "DELETE", Path: "/roles", Handler: handlers.DeleteRoleHandler},
}
