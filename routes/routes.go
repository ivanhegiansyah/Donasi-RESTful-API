package routes

import (

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo{
	e := echo.New()
	// g := e.Group("api/v1/")
	UserRoute(e.Group("api/v1/"))
	
	return e
}