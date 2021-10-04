package transactions

import (
	"context"
	"errors"
	"time"
)

type TransactionUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewTransactionUsecase(repo Repository, timeout time.Duration) Usecase {
	return &TransactionUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisnis tambah penggalangan dana
func (uc *TransactionUsecase) AddTransaction(ctx context.Context, domain Domain) (Domain, error) {
	if domain.TotalDonation == 0 {
		return Domain{}, errors.New("Nominal empty")
	}
	transaction, err := uc.Repo.AddTransaction(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return transaction, nil
}

//core bisnis Read
func (uc *TransactionUsecase) GetAllTransaction(ctx context.Context) ([]Domain, error) {
	transaction, err := uc.Repo.GetAllTransaction(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return transaction, nil
}