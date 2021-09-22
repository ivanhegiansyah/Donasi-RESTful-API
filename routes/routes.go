package routes

import (
	"finalproject-BE/controllers"
	"finalproject-BE/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	e.Pre(middleware.RemoveTrailingSlash())
	ev1.POST("users/register", controllers.RegisterUserController)
	ev1.POST("users/login", controllers.LoginUserController)
	MiddlewareRoute(ev1)

	middlewares.LogMiddleware(e)

	return e
}
