package routes

import (
	donationdetails "finalproject-BE/controllers/donationDetails"
	donationtypes "finalproject-BE/controllers/donationTypes"
	"finalproject-BE/controllers/donations"
	"finalproject-BE/controllers/news"
	paymentmethods "finalproject-BE/controllers/paymentMethods"
	"finalproject-BE/controllers/transactions"
	"finalproject-BE/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware            middleware.JWTConfig
	UserController           users.UserController
	DonationController       donations.DonationController
	DonationDetailController donationdetails.DonationDetailController
	DonationTypeController   donationtypes.DonationTypeController
	TransactionController    transactions.TransactionController
	PaymentMethodController  paymentmethods.PaymentMethodController
	NewsController           news.NewsController
}

func (cl *ControllerList) RouteUser(e *echo.Echo) {
	e.POST("/api/v1/users/login", cl.UserController.Login)
	e.POST("/api/v1/users/register", cl.UserController.Register)
	e.GET("/api/v1/users", cl.UserController.GetAllUser)
	e.GET("/api/v1/users/:id", cl.UserController.GetDetailUser, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/api/v1/users/:id", cl.UserController.UpdateUser, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.DELETE("/api/v1/users/:id", cl.UserController.DeleteUser, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func (cl *ControllerList) RouteDonation(e *echo.Echo) {

	e.POST("/api/v1/donations/add-donation", cl.DonationController.AddDonation, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/donations", cl.DonationController.GetAllDonation, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/donations/:id", cl.DonationController.GetDetailDonation, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/api/v1/donations/:id", cl.DonationController.Updatedonation, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.DELETE("/api/v1/donations/:id", cl.DonationController.DeleteDonation, middleware.JWTWithConfig(cl.JWTMiddleware))

}

func (cl *ControllerList) RouteDonationDetail(e *echo.Echo) {
	e.POST("/api/v1/donations/:id/add-detail", cl.DonationDetailController.AddDonationDetail, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/api/v1/donations/:id/update-detail", cl.DonationDetailController.UpdatedonationDetail, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.DELETE("/api/v1/donations/:id/delete-detail", cl.DonationDetailController.DeleteDonationDetail, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func (cl *ControllerList) RouteDonationType(e *echo.Echo) {
	e.POST("/api/v1/donations/add-type", cl.DonationTypeController.AddDonationType, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/donations/types", cl.DonationTypeController.GetDonationType, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/donations/types/:id", cl.DonationTypeController.GetDetailDonationType, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.PUT("/api/v1/donations/types/:id", cl.DonationTypeController.UpdateDonationType, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.DELETE("/api/v1/donations/types/:id", cl.DonationTypeController.DeleteDonationType, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func (cl *ControllerList) RouteTransaction(e *echo.Echo) {
	e.POST("/api/v1/transactions/add-transaction", cl.TransactionController.AddTransaction, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/transactions", cl.TransactionController.GetAllTransaction, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func (cl *ControllerList) RoutePaymentMethod(e *echo.Echo) {
	e.POST("/api/v1/transactions/add-payment", cl.PaymentMethodController.AddPaymentMethod, middleware.JWTWithConfig(cl.JWTMiddleware))
	e.GET("/api/v1/transactions/payment-method", cl.PaymentMethodController.GetAllPaymentMethod, middleware.JWTWithConfig(cl.JWTMiddleware))
}

func (cl *ControllerList) RouteNews(e *echo.Echo) {
	e.GET("/api/v1/news", cl.NewsController.GetByCategory)
}
