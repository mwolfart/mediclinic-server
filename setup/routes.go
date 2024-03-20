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
	{Method: "GET", Path: "/insurances", Handler: handlers.GetInsuranceHandler},
	{Method: "POST", Path: "/insurances", Handler: handlers.AddInsuranceHandler},
	{Method: "DELETE", Path: "/insurances/:id", Handler: handlers.DeleteInsuranceHandler},
	{Method: "GET", Path: "/roles", Handler: handlers.GetRoleHandler},
	{Method: "POST", Path: "/roles", Handler: handlers.AddRoleHandler},
	{Method: "DELETE", Path: "/roles/:id", Handler: handlers.DeleteRoleHandler},
}
