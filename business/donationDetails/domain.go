package donationdetails

import (
	"context"
	"time"
)

type Domain struct {
	Id          int
	DonationId  int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	AddDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error)
	UpdateDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error)
	DeleteDonationDetail(ctx context.Context, id int) error
}

type Repository interface {
	AddDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error)
	UpdateDonationDetail(ctx context.Context, domain Domain) (Domain, error)
	DeleteDonationDetail(ctx context.Context, id int) error
}
