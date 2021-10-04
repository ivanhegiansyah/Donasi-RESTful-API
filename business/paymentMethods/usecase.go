package paymentmethods

import (
	"context"
	"errors"
	"time"
)

type PaymentMethodUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewPaymentMethodUsecase(repo Repository, timeout time.Duration) Usecase {
	return &PaymentMethodUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisnis tambah penggalangan dana
func (uc *PaymentMethodUsecase) AddPaymentMethod(ctx context.Context, domain Domain) (Domain, error) {
	if domain.MethodName == "" {
		return Domain{}, errors.New("Payment empty")
	}
	payment, err := uc.Repo.AddPaymentMethod(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return payment, nil
}

//core bisnis Read
func (uc *PaymentMethodUsecase) GetAllPaymentMethod(ctx context.Context) ([]Domain, error) {
	payment, err := uc.Repo.GetAllPaymentMethod(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return payment, nil
}

