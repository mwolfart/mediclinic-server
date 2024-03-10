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
	&Route{Method: "GET", Path: "/users", Handler: handlers.GetUserHandler},
}
