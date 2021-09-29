package donationdetails

import (
	donationdetails "finalproject-BE/business/donationDetails"
	"time"

	"gorm.io/gorm"
)

type DonationDetails struct {
	Id          int `gorm:"primaryKey"`
	DonationId  int `gorm:"unique"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (donationDetail *DonationDetails) ToDomain() donationdetails.Domain {
	return donationdetails.Domain{
		Id:          donationDetail.Id,
		DonationId:  donationDetail.DonationId,
		Description: donationDetail.Description,
		CreatedAt:   donationDetail.CreatedAt,
		UpdatedAt:   donationDetail.UpdatedAt,
	}
}

func FromDomain(domain donationdetails.Domain) DonationDetails {
	return DonationDetails{
		Id:          domain.Id,
		DonationId:  domain.DonationId,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
