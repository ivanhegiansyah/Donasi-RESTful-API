package requests

import (
	"finalproject-BE/business/transactions"
)

type TransactionAdd struct {
	UserId          int    `json:"userId"`
	PaymentMethodId int    `json:"paymentMethodId"`
	DonationId      int    `json:"donationId"`
	TotalDonation   int    `json:"totalDonation"`
	Status          string `json:"status"`
}

func (transaction *TransactionAdd) ToDomainAdd() transactions.Domain {
	return transactions.Domain{
		UserId:          transaction.UserId,
		PaymentMethodId: transaction.PaymentMethodId,
		DonationId:      transaction.DonationId,
		TotalDonation:   transaction.TotalDonation,
		Status:          transaction.Status,
	}
}
