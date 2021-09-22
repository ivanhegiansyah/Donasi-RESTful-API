package responses

import (
	"finalproject-BE/business/users"
	"time"
)

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Dob       string    `json:"dob"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone:     domain.Phone,
		Dob:       domain.Dob,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
