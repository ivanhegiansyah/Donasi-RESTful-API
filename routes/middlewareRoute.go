package routes

import (
	"finalproject-BE/constants"
	"finalproject-BE/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddlewareRoute(g *echo.Group) {
	g.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	g.GET("users/:id", controllers.GetOneUserController)
	DonationRoute(g)
}
