package responses

import (
	"finalproject-BE/business/users"
	"finalproject-BE/drivers/databases/donations"
	"finalproject-BE/drivers/databases/transactions"
)

type UserDetailResponse struct {
	UserResponses UserResponse
	Transaction   []transactions.Transactions `json:"transaction"`
	Donation      []donations.Donations       `json:"donation"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Dob   string `json:"dob"`
}

func FromDomainDetail(domain users.Domain) UserDetailResponse {
	return UserDetailResponse{
		UserResponses: UserResponse{
			Id:    domain.Id,
			Name:  domain.Name,
			Email: domain.Email,
			Phone: domain.Phone,
			Dob:   domain.Dob},
		Transaction: domain.Transaction,
		Donation:    domain.Donation,
	}
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:    domain.Id,
		Name:  domain.Name,
		Email: domain.Email,
		Phone: domain.Phone,
		Dob:   domain.Dob,
	}
}
func FromUserListDomain(domain []users.Domain) []UserResponse {
	var response []UserResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}
