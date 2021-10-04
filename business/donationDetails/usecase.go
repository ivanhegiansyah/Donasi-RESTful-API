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

	domain.DonationId = id
	donationDetail, err := uc.Repo.AddDonationDetail(ctx, domain, id)

	if err != nil {
		return Domain{}, err
	}
	return donationDetail, nil
}

func (uc *DonationDetailUsecase) UpdateDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if domain.Description == "" {
		return Domain{}, errors.New("Description empty")
	}

	domain.DonationId = id

	donation, err := uc.Repo.UpdateDonationDetail(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return donation, nil
}

func (uc *DonationDetailUsecase) DeleteDonationDetail(ctx context.Context, id int) error {
	err := uc.Repo.DeleteDonationDetail(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
