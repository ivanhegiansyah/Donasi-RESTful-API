package paymentmethods

import (
	paymentmethods "finalproject-BE/business/paymentMethods"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/paymentMethods/requests"
	"finalproject-BE/controllers/paymentMethods/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentMethodController struct {
	PaymentMethodUseCase paymentmethods.Usecase
}

func NewPaymentMethodController(paymentMethodUseCase paymentmethods.Usecase) *PaymentMethodController {
	return &PaymentMethodController{
		PaymentMethodUseCase: paymentMethodUseCase,
	}
}

func (paymentMethodController PaymentMethodController) AddPaymentMethod(c echo.Context) error {
	fmt.Println("Add")
	paymentMethodadd := requests.PaymentMethodAdd{}
	c.Bind(&paymentMethodadd)
	ctx := c.Request().Context()
	paymentMethod, error := paymentMethodController.PaymentMethodUseCase.AddPaymentMethod(ctx, paymentMethodadd.ToDomainAdd())
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(paymentMethod))

}

func (paymentMethodController PaymentMethodController) GetAllPaymentMethod(c echo.Context) error {
	fmt.Println("GetAll")
	ctx := c.Request().Context()
	paymentMethod, error := paymentMethodController.PaymentMethodUseCase.GetAllPaymentMethod(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, paymentMethod)
}
