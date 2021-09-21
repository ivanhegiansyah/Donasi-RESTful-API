package routes

import (
	"finalproject-BE/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(g *echo.Group) {
	g.GET("users", controllers.GetAllUserController) //blm tambahin query param
	g.POST("users/register", controllers.RegisterUserController)
	g.POST("users/login", controllers.LoginUserController)
	g.PUT("users/:id", controllers.UpdateUserController)
	g.DELETE("users/:id", controllers.DeleteUserController)
}
