package responses

import (
	"finalproject-BE/business/transactions"
	"time"
)

type TransactionResponse struct {
	Id              int       `json:"id"`
	UserId          int       `json:"userId"`
	PaymentMethodId int       `json:"paymentMethodId"`
	DonationId      int       `json:"donationId"`
	TotalDonation   int       `json:"totalDonation"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func FromDomain(domain transactions.Domain) TransactionResponse {
	return TransactionResponse{
		Id:              domain.Id,
		UserId:          domain.UserId,
		PaymentMethodId: domain.PaymentMethodId,
		DonationId:      domain.DonationId,
		TotalDonation: domain.TotalDonation,
		Status:          domain.Status,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}
