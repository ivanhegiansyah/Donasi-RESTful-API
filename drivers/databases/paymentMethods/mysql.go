package paymentmethods

import (
	"context"
	paymentmethods "finalproject-BE/business/paymentMethods"

	"gorm.io/gorm"
)

type MysqlPaymentMethodRepository struct {
	Conn *gorm.DB
}

func NewMysqlPaymentMethodRepository(conn *gorm.DB) paymentmethods.Repository {
	return &MysqlPaymentMethodRepository{
		Conn: conn,
	}
}

func (rep *MysqlPaymentMethodRepository) AddPaymentMethod(ctx context.Context, domain paymentmethods.Domain) (paymentmethods.Domain, error) {
	var paymentMethod PaymentMethods

	paymentMethod = FromDomain(domain)

	result := rep.Conn.Create(&paymentMethod)

	if result.Error != nil {
		return paymentmethods.Domain{}, result.Error
	}

	return paymentMethod.ToDomain(), nil
}

func (rep *MysqlPaymentMethodRepository) GetAllPaymentMethod(ctx context.Context) ([]paymentmethods.Domain, error) {
	var paymentMethod []PaymentMethods

	result := rep.Conn.Find(&paymentMethod)

	if result.Error != nil {
		return []paymentmethods.Domain{}, result.Error
	}

	return ToListDomain(paymentMethod), nil
}
