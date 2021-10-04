package responses

import (
	paymentmethods "finalproject-BE/business/paymentMethods"
	"time"
)

type PaymentMethodResponse struct {
	Id         int       `json:"id"`
	MethodName string    `json:"methodName"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}


func FromDomain(domain paymentmethods.Domain) PaymentMethodResponse {
	return PaymentMethodResponse{
		Id:         domain.Id,
		MethodName: domain.MethodName,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromListDomain(domain []paymentmethods.Domain) []PaymentMethodResponse {
	var response []PaymentMethodResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}