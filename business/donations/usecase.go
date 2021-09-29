package donations

import (
	"context"
	"errors"
	"time"
)

type DonationUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewDonationUsecase(repo Repository, timeout time.Duration) Usecase {
	return &DonationUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

// core bisnis tambah penggalangan dana
func (uc *DonationUsecase) AddDonation(ctx context.Context, domain Domain) (Domain, error) {
	if domain.DonationName == "" {
		return Domain{}, errors.New("Name empty")
	}
	if domain.ShortDescription == "" {
		return Domain{}, errors.New("Short Description empty")
	}
	if domain.GoalAmount == 0 {
		return Domain{}, errors.New("Goal amount empty")
	}
	
	donation, err := uc.Repo.AddDonation(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return donation, nil
}

//core bisnis Read
func (uc *DonationUsecase) GetAllDonation(ctx context.Context) ([]Domain, error) {
	donation, err := uc.Repo.GetAllDonation(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return donation, nil
}

func (uc *DonationUsecase) GetDetailDonation(ctx context.Context, id int) (Domain, error) {
	donation, err := uc.Repo.GetDetailDonation(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return donation, nil
}

func (uc *DonationUsecase) UpdateDonation(ctx context.Context, domain Domain, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if domain.DonationName == "" {
		return Domain{}, errors.New("Name empty")
	}
	if domain.ShortDescription == "" {
		return Domain{}, errors.New("Short Description empty")
	}
	if domain.GoalAmount == 0 {
		return Domain{}, errors.New("Goal amount empty")
	}

	domain.Id = id

	donation, err := uc.Repo.UpdateDonation(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return donation, nil
}

func (uc *DonationUsecase) DeleteDonation(ctx context.Context, id int) error {
	err := uc.Repo.DeletDonation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}