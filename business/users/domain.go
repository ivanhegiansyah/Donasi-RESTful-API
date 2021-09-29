package users

import (
	"context"
	"finalproject-BE/drivers/databases/donations"
	"finalproject-BE/drivers/databases/transactions"
	"time"
)

// type DetailDonate struct {
// 	ID               int
// 	DonationName     string
// 	Status           string
// 	ShortDescription string
// 	GoalAmount       int
// 	ExpiredDate      string
// }

type Domain struct {
	Id        int
	Name      string
	Email     string
	Password  string
	Phone     string
	Dob       string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Transaction []transactions.Transactions
	Donation  []donations.Donations
}

type Usecase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	GetDetailUser(ctx context.Context, id int) (Domain, error)
	UpdateUser(ctx context.Context, domain Domain, id int) (Domain, error)
	DeleteUser(ctx context.Context, id int) error
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	GetDetailUser(ctx context.Context, id int) (Domain, error)
	UpdateUser(ctx context.Context, domain Domain) (Domain, error)
	DeleteUser(ctx context.Context, id int) error
}
