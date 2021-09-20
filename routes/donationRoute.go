package routes

import (
	"finalproject-BE/controllers"

	"github.com/labstack/echo/v4"
)

func DonationRoute(g *echo.Group) {
	g.GET("users/:userid/donations", controllers.GetAllDonationController) //blm tambahin query param
	g.GET("users/:userid/donations/:donationid", controllers.GetOneDonationController)
	g.POST("users/:userid/add-donation", controllers.AddDonationrController)
	g.PUT("users/:userid/donations/:donationid", controllers.UpdateDonationController)
	g.DELETE("users/:userid/donations/:donationid", controllers.DeleteDonationController)
}
