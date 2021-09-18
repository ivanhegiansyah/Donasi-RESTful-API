package routes

import (
	"finalproject-BE/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(g *echo.Group) {
	g.GET("users", controllers.GetAllUserController)
	g.GET("users/:id", controllers.GetOneUserController)
	g.POST("users/register", controllers.RegisterUserController)
	g.PUT("users/:id", controllers.UpdateUserController)
	g.DELETE("users/:id", controllers.DeleteUserController)
}
