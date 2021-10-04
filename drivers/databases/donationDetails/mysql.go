package donationdetails

import (
	"context"
	"errors"
	donationdetails "finalproject-BE/business/donationDetails"

	"gorm.io/gorm"
)

type MysqlDonationDetailRepository struct {
	Conn *gorm.DB
}

func NewMysqlDonationDetailRepository(conn *gorm.DB) donationdetails.Repository {
	return &MysqlDonationDetailRepository{
		Conn: conn,
	}
}

func (rep *MysqlDonationDetailRepository) AddDonationDetail(ctx context.Context, domain donationdetails.Domain, id int) (donationdetails.Domain, error) {
	var donationDetail DonationDetails

	donationDetail = FromDomain(domain)

	result := rep.Conn.Create(&donationDetail)

	if result.Error != nil {
		return donationdetails.Domain{}, result.Error
	}

	return donationDetail.ToDomain(), nil
}

func (rep *MysqlDonationDetailRepository) UpdateDonationDetail(ctx context.Context, domain donationdetails.Domain) (donationdetails.Domain, error) {
	var donationDetail DonationDetails
	var detailId int

	donationDetail = FromDomain(domain)

	rep.Conn.Raw("SELECT id FROM donation_details").Scan(&detailId)
	donationDetail.Id = detailId
	result := rep.Conn.Save(&donationDetail)

	if result.Error != nil {
		return donationdetails.Domain{}, result.Error
	}

	return donationDetail.ToDomain(), nil

}

func (rep *MysqlDonationDetailRepository) DeleteDonationDetail(ctx context.Context, id int) error {
	var donationDetail DonationDetails
	var detailId int

	rep.Conn.Raw("SELECT id FROM donation_details").Scan(&detailId)
	result := rep.Conn.Delete(&donationDetail, detailId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}