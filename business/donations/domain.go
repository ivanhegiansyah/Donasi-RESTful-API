package donations

import (
	"context"
	donationdetails "finalproject-BE/drivers/databases/donationDetails"
	"time"
)

type Domain struct {
	Id               int
	UserId           int
	DonationTypeId   int
	DonationName     string
	Status           string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	ExpiredDate      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DonationDetail   donationdetails.DonationDetails
}

type Usecase interface {
	AddDonation(ctx context.Context, domain Domain) (Domain, error)
	GetAllDonation(ctx context.Context) ([]Domain, error)
	GetDetailDonation(ctx context.Context, id int) (Domain, error)
	UpdateDonation(ctx context.Context, domain Domain, id int) (Domain, error)
	DeleteDonation(ctx context.Context, id int) error
}

type Repository interface {
	AddDonation(ctx context.Context, domain Domain) (Domain, error)
	GetAllDonation(ctx context.Context) ([]Domain, error)
	GetDetailDonation(ctx context.Context, id int) (Domain, error)
	UpdateDonation(ctx context.Context, domain Domain) (Domain, error)
	DeleteDonation(ctx context.Context, id int) error
}
