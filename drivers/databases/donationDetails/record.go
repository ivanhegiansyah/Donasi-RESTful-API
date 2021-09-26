package donationdetails

import (
	donationdetails "finalproject-BE/business/donationDetails"
	"finalproject-BE/drivers/databases/donations"
	"time"

	"gorm.io/gorm"
)

type DonationDetails struct {
	Id          int `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt      `gorm:"index"`
	Donation    donations.Donations `gorm:"foreignKey:DonationDetailId;references:Id"`
}

func (donationDetail *DonationDetails) ToDomain() donationdetails.Domain {
	return donationdetails.Domain{
		Id:          donationDetail.Id,
		Description: donationDetail.Description,
		CreatedAt:   donationDetail.CreatedAt,
		UpdatedAt:   donationDetail.UpdatedAt,
	}
}

func FromDomain(domain donationdetails.Domain) DonationDetails {
	return DonationDetails{
		Id:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
