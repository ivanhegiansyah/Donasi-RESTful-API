package routes

import (
	"finalproject-BE/controllers/donations"
	"finalproject-BE/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
	DonationController donations.DonationController
}

func (cl *ControllerList) RouteUser (e *echo.Echo) {
	e.POST("/api/v1/users/login", cl.UserController.Login)
	e.POST("/api/v1/users/register", cl.UserController.Register)
	e.GET("/api/v1/users", cl.UserController.GetAllUser)
	e.GET("/api/v1/users/:id", cl.UserController.GetDetailUser)
}

func (cl *ControllerList) RouteDonation (e *echo.Echo) {
	e.POST("/api/v1/donations/add-donation", cl.DonationController.AddDonation)
	e.GET("/api/v1/donations", cl.DonationController.GetAllDonation)
	e.GET("/api/v1/donations/:id", cl.DonationController.GetDetailDonation)
}

