package donationtypes

import (
	"context"
	donationtypes "finalproject-BE/business/donationTypes"

	"gorm.io/gorm"
)

type MysqlDonationTypeRepository struct {
	Conn *gorm.DB
}

func NewMysqlDonationTypeRepository(conn *gorm.DB) donationtypes.Repository {
	return &MysqlDonationTypeRepository{
		Conn: conn,
	}
}

func (rep *MysqlDonationTypeRepository) AddDonationType(ctx context.Context, domain donationtypes.Domain, id int) (donationtypes.Domain, error) {
	var donationType DonationTypes

	donationType = FromDomain(domain)

	result := rep.Conn.Create(&donationType)

	if result.Error != nil {
		return donationtypes.Domain{}, result.Error
	}

	return donationType.ToDomain(), nil
}
