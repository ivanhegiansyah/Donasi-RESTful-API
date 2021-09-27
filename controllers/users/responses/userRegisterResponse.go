package responses

import (
	"finalproject-BE/business/users"
	"time"
)

type UserRegisterResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Dob       string    `json:"dob"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomainRegister(domain users.Domain) UserRegisterResponse {
	return UserRegisterResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		Email:     domain.Email,
		Phone:     domain.Phone,
		Dob:       domain.Dob,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
