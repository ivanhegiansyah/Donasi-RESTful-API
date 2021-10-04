package transactions_test

import (
	"context"
	"finalproject-BE/business/transactions"
	mockTransaction "finalproject-BE/drivers/databases/transactions/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	transactionRepository mockTransaction.TransactionRepository
	transactionService    transactions.Usecase
	transactionDomain     transactions.Domain
)

func setup() {
	transactionService = transactions.NewTransactionUsecase(&transactionRepository, time.Hour*1)
	transactionDomain = transactions.Domain{
		Id:               1,
		UserId:           1,
		PaymentMethodId:   1,
		DonationId:     1,
		TotalDonation: 250000,
		Status:           "New",
	}
}

func TestAddTransaction(t *testing.T) {
	setup()
	transactionRepository.On("AddTransaction",
		mock.Anything,
		mock.AnythingOfType("Domain")).Return(transactionDomain, nil).Once()
	t.Run("Test Case 1 | Valid Add Transaction", func(t *testing.T) {
		_, err := transactionService.AddTransaction(context.Background(), transactionDomain)

		assert.Nil(t, err)

	})
	t.Run("Test Case 2 | Invalid Add Transaction Field Empty", func(t *testing.T) {
		_, err := transactionService.AddTransaction(context.Background(), transactions.Domain{
			Id:               1,
		UserId:           1,
		PaymentMethodId:   1,
		DonationId:     1,
		TotalDonation: 0,
		Status:           "New",
		})
		assert.NotNil(t, err)
	})
}

func TestGetAllTransaction(t *testing.T) {
	setup()
	transactionRepository.On("GetAllTransaction",
		mock.Anything).Return([]transactions.Domain{}, nil).Once()
	t.Run("Test Case 1 | Valid Get All Donation", func(t *testing.T) {
		_, err := transactionService.GetAllTransaction(context.Background())

		assert.Nil(t, err)
	})
}

