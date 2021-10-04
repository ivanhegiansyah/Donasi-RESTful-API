package transactions

import (
	"finalproject-BE/business/transactions"
	"finalproject-BE/controllers"
	"finalproject-BE/controllers/transactions/requests"
	"finalproject-BE/controllers/transactions/responses"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionUseCase transactions.Usecase
}

func NewTransactionController(transactionUseCase transactions.Usecase) *TransactionController {
	return &TransactionController{
		TransactionUseCase: transactionUseCase,
	}
}

func (transactionController TransactionController) AddTransaction(c echo.Context) error {
	fmt.Println("Add")
	transactionadd := requests.TransactionAdd{}
	c.Bind(&transactionadd)
	ctx := c.Request().Context()
	transaction, error := transactionController.TransactionUseCase.AddTransaction(ctx, transactionadd.ToDomainAdd())
	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(transaction))

}

func (transactionController TransactionController) GetAllTransaction(c echo.Context) error {
	fmt.Println("GetAll")
	ctx := c.Request().Context()
	transaction, error := transactionController.TransactionUseCase.GetAllTransaction(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(transaction))
}

