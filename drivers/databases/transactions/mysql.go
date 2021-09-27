package transactions

import (
	"context"
	"finalproject-BE/business/transactions"
	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlTransactionRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (rep *MysqlTransactionRepository) AddTransaction(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	var transaction Transactions
	
	transaction = FromDomain(domain)
	result := rep.Conn.Create(&transaction)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	return transaction.ToDomain(), nil
}

func (rep *MysqlTransactionRepository) GetAllTransaction(ctx context.Context) ([]transactions.Domain, error) {
	var transaciton []Transactions

	result := rep.Conn.Find(&transaciton)

	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}

	return ToListDomain(transaciton), nil
}

func (rep *MysqlTransactionRepository) GetDetailTransaction(ctx context.Context, id int) ([]transactions.Domain, error) {
	var transaction []Transactions

	result := rep.Conn.First(&transaction, id)
	
	

	if result.Error != nil {
		return []transactions.Domain{}, result.Error
	}

	return ToListDomain(transaction), nil
}