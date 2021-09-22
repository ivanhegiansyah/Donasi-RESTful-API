package users

import (
	"context"
	"finalproject-BE/controllers/users/requests"
	"time"
)

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
}

type Usecase interface {
	Register(ctx context.Context, userRequest requests.UserRegister) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, userRequest requests.UserRegister) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}
