package transactions

import (
	"finalproject-BE/business/transactions"
	"finalproject-BE/drivers/databases/donations"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	Id              int `gorm:"primaryKey"`
	UserId          int
	PaymentMethodId int
	DonationId      int
	TotalDonation   int
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt        `gorm:"index"`
	Donations       []donations.Donations `gorm:"many2many:transactions_donations;"`
}

func (transaction *Transactions) ToDomain() transactions.Domain {
	return transactions.Domain{
		Id:              transaction.Id,
		UserId:          transaction.UserId,
		PaymentMethodId: transaction.PaymentMethodId,
		DonationId:      transaction.DonationId,
		TotalDonation:   transaction.TotalDonation,
		Status:          transaction.Status,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		Id:              domain.Id,
		UserId:          domain.UserId,
		PaymentMethodId: domain.PaymentMethodId,
		DonationId:      domain.DonationId,
		TotalDonation:   domain.TotalDonation,
		Status:          "New",
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}

func ToListDomain(data []Transactions) []transactions.Domain {
	result := []transactions.Domain{}
	for _, domain := range data {
		result = append(result, domain.ToDomain())
	}
	return result
}
