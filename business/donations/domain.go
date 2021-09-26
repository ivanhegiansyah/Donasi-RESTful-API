package donations

import (
	"context"
	"time"
)

type Domain struct {
	Id               int
	UserId           int
	DonationDetailId int
	DonationTypeId   int
	DonationName     string
	Status           string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	ExpiredDate      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	AddDonation(ctx context.Context, domain Domain) (Domain, error)
	GetAllDonation(ctx context.Context) ([]Domain, error)
	GetDetailDonation(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	AddDonation(ctx context.Context, domain Domain) (Domain, error)
	GetAllDonation(ctx context.Context) ([]Domain, error)
	GetDetailDonation(ctx context.Context, id int) ([]Domain, error)
}
