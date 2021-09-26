package donationdetails

import (
	"context"
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
