package donationtypes

import (
	"context"
	"finalproject-BE/drivers/databases/donations"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Donations  []donations.Donations
}

type Usecase interface {
	AddDonationType(ctx context.Context, domain Domain) (Domain, error)
	GetDonationType(ctx context.Context) ([]Domain, error)
	GetDetailDonationType(ctx context.Context, id int) (Domain, error)
	UpdateDonationType(ctx context.Context, domain Domain, id int) (Domain, error)
	DeleteDonationType(ctx context.Context, id int) error
}

type Repository interface {
	AddDonationType(ctx context.Context, domain Domain) (Domain, error)
	GetDonationType(ctx context.Context) ([]Domain, error)
	GetDetailDonationType(ctx context.Context, id int) (Domain, error)
	UpdateDonationType(ctx context.Context, domain Domain) (Domain, error)
	DeleteDonationType(ctx context.Context, id int) error
}
