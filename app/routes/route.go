package routes

import (
	"finalproject-BE/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister (e *echo.Echo) {
	e.POST("/api/v1/users/login", cl.UserController.Login)
	e.POST("/api/v1/users/register", cl.UserController.Register)
}