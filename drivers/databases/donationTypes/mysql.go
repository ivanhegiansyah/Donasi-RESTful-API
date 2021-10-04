package donationtypes

import (
	"context"
	"errors"
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

func (rep *MysqlDonationTypeRepository) AddDonationType(ctx context.Context, domain donationtypes.Domain) (donationtypes.Domain, error) {
	var donationType DonationTypes

	donationType = FromDomain(domain)

	result := rep.Conn.Create(&donationType)

	if result.Error != nil {
		return donationtypes.Domain{}, result.Error
	}

	return donationType.ToDomain(), nil
}

func (rep *MysqlDonationTypeRepository) GetDonationType(ctx context.Context) ([]donationtypes.Domain, error) {
	var donationType []DonationTypes

	result := rep.Conn.Find(&donationType)

	if result.Error != nil {
		return []donationtypes.Domain{}, result.Error
	}

	return ToListDomain(donationType), nil
}

func (rep *MysqlDonationTypeRepository) GetDetailDonationType(ctx context.Context, id int) (donationtypes.Domain, error) {
	var donationType DonationTypes

	result := rep.Conn.Preload("Donations").First(&donationType, id)

	if result.Error != nil {
		return donationtypes.Domain{}, result.Error
	}

	return donationType.ToDomain(), nil
}

func (rep *MysqlDonationTypeRepository) UpdateDonationType(ctx context.Context, domain donationtypes.Domain) (donationtypes.Domain, error) {
	var donationType DonationTypes

	donationType = FromDomain(domain)

	result := rep.Conn.Save(&donationType)

	if result.Error != nil {
		return donationtypes.Domain{}, result.Error
	}

	return donationType.ToDomain(), nil

}

func (rep *MysqlDonationTypeRepository) DeleteDonationType(ctx context.Context, id int) error {
	var donationType DonationTypes

	result := rep.Conn.Delete(&donationType, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
