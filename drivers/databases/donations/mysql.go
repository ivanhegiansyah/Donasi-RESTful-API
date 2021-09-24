package donations

import (
	"context"
	"finalproject-BE/business/donations"

	"gorm.io/gorm"
)

type MysqlDonationRepository struct {
	Conn *gorm.DB
}

func NewMysqlDonationRepository(conn *gorm.DB) donations.Repository {
	return &MysqlDonationRepository{
		Conn: conn,
	}
}

func (rep *MysqlDonationRepository) AddDonation(ctx context.Context, domain donations.Domain) (donations.Domain, error) {
	var donation Donations
	
	donation.DonationName = domain.DonationName
	donation.Status = domain.Status
	donation.ShortDescription = domain.ShortDescription
	donation.GoalAmount = domain.GoalAmount
	donation.CurrentAmount = domain.CurrentAmount
	donation.ExpiredDate = domain.ExpiredDate

	result := rep.Conn.Create(&donation)

	if result.Error != nil {
		return donations.Domain{}, result.Error
	}

	return donation.ToDomain(), nil
}

func (rep *MysqlDonationRepository) GetAllDonation(ctx context.Context) ([]donations.Domain, error) {
	var user []Donations

	result := rep.Conn.Find(&user)

	if result.Error != nil {
		return []donations.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}

func (rep *MysqlDonationRepository) GetDetailDonation(ctx context.Context, id int) ([]donations.Domain, error) {
	var donation []Donations

	result := rep.Conn.First(&donation, id)

	if result.Error != nil {
		return []donations.Domain{}, result.Error
	}

	return ToListDomain(donation), nil
}