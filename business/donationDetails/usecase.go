package donationdetails

import (
	"context"
	"errors"
	"time"
)

type DonationDetailUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewDonationDetailUsecase(repo Repository, timeout time.Duration) Usecase {
	return &DonationDetailUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *DonationDetailUsecase) AddDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error) {
	if domain.Description == "" {
		return Domain{}, errors.New("Description empty")
	}
	donationDetail, err := uc.Repo.AddDonationDetail(ctx, domain, id)

	if err != nil {
		return Domain{}, err
	}
	return donationDetail, nil
}
