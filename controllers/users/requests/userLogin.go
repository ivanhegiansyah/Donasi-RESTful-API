package requests

import "finalproject-BE/business/users"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToDomainLogin() users.Domain {
	return users.Domain{
		Email:    user.Email,
		Password: user.Password,
	}
}