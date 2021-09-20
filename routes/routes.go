package routes

import (
	"finalproject-BE/middlewares"

	"github.com/labstack/echo/v4"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	// g := e.Group("api/v1/")
	UserRoute(e.Group("api/v1/"))
	MiddlewareRoute(e.Group("api/v1/"))

	middlewares.LogMiddleware(e)

	return e
}
