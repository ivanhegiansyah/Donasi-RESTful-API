package requests

import "finalproject-BE/business/users"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
}

func (user *UserRegister) ToDomainRegister() users.Domain {
	return users.Domain{
		Name: user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone: user.Phone,
		Dob: user.Dob,
	}
}