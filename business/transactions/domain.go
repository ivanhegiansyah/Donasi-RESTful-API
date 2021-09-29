package transactions

import (
	"context"
	"time"
)

type Domain struct {
	Id              int
	UserId          int
	PaymentMethodId int
	DonationId      int
	TotalDonation   int
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	
}

type Usecase interface {
	AddTransaction(ctx context.Context, domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	AddTransaction(ctx context.Context, domain Domain) (Domain, error)
	GetAllTransaction(ctx context.Context) ([]Domain, error)
}
