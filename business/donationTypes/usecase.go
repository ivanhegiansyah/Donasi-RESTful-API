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

func (uc *DonationTypeUsecase) AddDonationType(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Type name empty")
	}
	donationType, err := uc.Repo.AddDonationType(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return donationType, nil
}

//core bisnis Read
func (uc *DonationTypeUsecase) GetDonationType(ctx context.Context) ([]Domain, error) {
	donationType, err := uc.Repo.GetDonationType(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return donationType, nil
}

func (uc *DonationTypeUsecase) GetDetailDonationType(ctx context.Context, id int) (Domain, error) {
	donationType, err := uc.Repo.GetDetailDonationType(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return donationType, nil
}

func (uc *DonationTypeUsecase) UpdateDonationType(ctx context.Context, domain Domain, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if domain.Name == "" {
		return Domain{}, errors.New("Type name empty")
	}

	domain.Id = id

	donation, err := uc.Repo.UpdateDonationType(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return donation, nil
}

func (uc *DonationTypeUsecase) DeleteDonationType(ctx context.Context, id int) error {
	err := uc.Repo.DeleteDonationType(ctx, id)
	if err != nil {
		return err
	}
	return nil
}