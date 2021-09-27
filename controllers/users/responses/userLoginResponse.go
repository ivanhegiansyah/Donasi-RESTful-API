package responses

import (
	"finalproject-BE/business/users"
)

type UserLoginResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

func FromDomainLogin(domain users.Domain) UserLoginResponse {
	return UserLoginResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Token:     domain.Token,
	}
}
