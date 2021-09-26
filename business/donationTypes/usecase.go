package donationtypes

import (
	"context"
	"errors"
	"time"
)

type DonationTypeUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewDonationTypeUsecase(repo Repository, timeout time.Duration) Usecase {
	return &DonationTypeUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *DonationTypeUsecase) AddDonationType(ctx context.Context, domain Domain, id int) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Type name empty")
	}
	donationType, err := uc.Repo.AddDonationType(ctx, domain, id)

	if err != nil {
		return Domain{}, err
	}
	return donationType, nil
}
