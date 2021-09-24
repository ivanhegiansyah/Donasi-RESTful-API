package users

import (
	"context"
	"errors"
	"finalproject-BE/helpers/encrypt"
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
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("Email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("Password empty")
	}

	user, err := uc.Repo.Login(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

// core bisnis register
func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("Name empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("Email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("Password empty")
	}

	var err error
	domain.Password, err = encrypt.Hash(domain.Password)

	if err != nil {
		return Domain{}, err
	}

	user, err := uc.Repo.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

//core bisnis Read
func (uc *UserUsecase) GetAllUser(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAllUser(ctx)

	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetDetailUser(ctx context.Context, id int) ([]Domain, error) {
	user, err := uc.Repo.GetDetailUser(ctx, id)

	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}