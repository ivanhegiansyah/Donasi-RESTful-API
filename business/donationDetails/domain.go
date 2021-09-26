package donationdetails

import (
	"context"
	"time"
)

type Domain struct {
	Id               int
	Description string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	AddDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	AddDonationDetail(ctx context.Context, domain Domain, id int) (Domain, error)
}