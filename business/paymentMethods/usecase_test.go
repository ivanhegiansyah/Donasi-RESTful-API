package paymentmethods_test

import (
	"context"
	paymentmethods "finalproject-BE/business/paymentMethods"
	mockPayment "finalproject-BE/drivers/databases/paymentMethods/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	paymentMethodRepository mockPayment.PaymentMethodRepository
	paymentMethodService    paymentmethods.Usecase
	paymentMethodDomain     paymentmethods.Domain
)

func setup() {
	paymentMethodService = paymentmethods.NewPaymentMethodUsecase(&paymentMethodRepository, time.Hour*1)
	paymentMethodDomain = paymentmethods.Domain{
		Id:         1,
		MethodName: "DANA",
	}
}

func TestAddPaymentMethod(t *testing.T) {
	setup()
	paymentMethodRepository.On("AddPaymentMethod",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(paymentMethodDomain, nil).Once()
	t.Run("Test Case 1 | Valid Add PaymentMethod", func(t *testing.T) {
		_, err := paymentMethodService.AddPaymentMethod(context.Background(), paymentMethodDomain)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Add PaymentMethod Field Empty", func(t *testing.T) {
		_, err := paymentMethodService.AddPaymentMethod(context.Background(), paymentmethods.Domain{
			Id:         1,
			MethodName: "",
		})
		assert.NotNil(t, err)
	})
}

func TestGetAllPaymentMethod(t *testing.T) {
	setup()
	paymentMethodRepository.On("GetAllPaymentMethod",
		mock.Anything).Return([]paymentmethods.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get All Donation", func(t *testing.T) {
		_, err := paymentMethodService.GetAllPaymentMethod(context.Background())

		assert.Nil(t, err)
	})
}
