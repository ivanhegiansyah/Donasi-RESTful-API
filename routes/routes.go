package routes

import (
	"finalproject-BE/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	UserRoute(e.Group("api/v1/"))
	MiddlewareRoute(e.Group("api/v1/"))

	middlewares.LogMiddleware(e)

	return e
}
