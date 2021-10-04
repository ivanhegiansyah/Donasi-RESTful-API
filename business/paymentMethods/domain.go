package paymentmethods

import (
	"context"
	"time"
	"finalproject-BE/drivers/databases/transactions"
)

type Domain struct {
	Id         int
	MethodName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Transaction   []transactions.Transactions
}

type Usecase interface {
	AddPaymentMethod(ctx context.Context, domain Domain) (Domain, error)
	GetAllPaymentMethod(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddPaymentMethod(ctx context.Context, domain Domain) (Domain, error)
	GetAllPaymentMethod(ctx context.Context) ([]Domain, error)
}
