package paymentmethods

import (
	paymentmethods "finalproject-BE/business/paymentMethods"
	"finalproject-BE/drivers/databases/transactions"
	"time"

	"gorm.io/gorm"
)

type PaymentMethods struct {
	Id          int    `gorm:"primaryKey"`
	MethodName  string `gorm:"unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt              `gorm:"index"`
	Transaction []transactions.Transactions `gorm:"foreignKey:UserId;references:Id"`
}

func (paymentMethod *PaymentMethods) ToDomain() paymentmethods.Domain {
	return paymentmethods.Domain{
		Id:          paymentMethod.Id,
		MethodName: paymentMethod.MethodName,
		CreatedAt:   paymentMethod.CreatedAt,
		UpdatedAt:   paymentMethod.UpdatedAt,
		Transaction: paymentMethod.Transaction,
	}
}

func FromDomain(domain  paymentmethods.Domain) PaymentMethods {
	return PaymentMethods{
		Id:        domain.Id,
		MethodName: domain.MethodName,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Transaction: domain.Transaction,
	}
}

func ToListDomain(data []PaymentMethods) []paymentmethods.Domain {
	result := []paymentmethods.Domain{}

	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return result
}
