package users

import (
	"context"
	"errors"
	"finalproject-BE/controllers/users/requests"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core bisnis login
func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("Email empty")
	}

	if password == "" {
		return Domain{}, errors.New("Password empty")
	}

	user, err := uc.Repo.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

//core bisnis register
func (uc *UserUsecase) Register(ctx context.Context, userRegister requests.UserRegister) (Domain, error) {
	if userRegister.Name == "" {
		return Domain{}, errors.New("Name empty")
	}
	if userRegister.Email == "" {
		return Domain{}, errors.New("Email empty")
	}
	if userRegister.Password == "" {
		return Domain{}, errors.New("Password empty")
	}
	user, err := uc.Repo.Register(ctx, userRegister)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
