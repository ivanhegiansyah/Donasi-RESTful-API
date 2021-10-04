package transactions

import (
	"context"
	"finalproject-BE/business/transactions"
	"finalproject-BE/drivers/databases/donations"

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

func (t *Transactions) AfterCreate(tx *gorm.DB) (err error) {
	donation := donations.Donations{}
	var current, goal int
	tx.Raw("SELECT current_amount FROM donations WHERE id = ?", t.DonationId).Scan(&current)
	tx.Raw("SELECT goal_amount FROM donations WHERE id = ?", t.DonationId).Scan(&goal)

	donation.CurrentAmount = current + t.TotalDonation
	tx.Model(&donation).Where("id = ?", t.DonationId).Update("current_amount", donation.CurrentAmount)
	
	if donation.CurrentAmount >= goal {
		tx.Model(&donation).Where("id = ?", t.DonationId).Update("status", "goal reached")
	}

	return
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
