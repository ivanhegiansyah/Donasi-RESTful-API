package users

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	Name         string
	Email        string
	Password     string
	Phone        string
	Dob          string
	Token        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	GetDetailUser(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetAllUser(ctx context.Context) ([]Domain, error)
	GetDetailUser(ctx context.Context, id int) (Domain, error)
}
