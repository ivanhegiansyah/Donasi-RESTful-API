package donationtypes

import (
	donationtypes "finalproject-BE/business/donationTypes"
	"finalproject-BE/drivers/databases/donations"
	"time"

	"gorm.io/gorm"
)

type DonationTypes struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt        `gorm:"index"`
	Donations  []donations.Donations `gorm:"foreignKey:DonationTypeId;references:Id"`
}

func (donationType *DonationTypes) ToDomain() donationtypes.Domain {
	return donationtypes.Domain{
		Id:        donationType.Id,
		Name:      donationType.Name,
		CreatedAt: donationType.CreatedAt,
		UpdatedAt: donationType.UpdatedAt,
		Donations:  donationType.Donations,
	}
}

func FromDomain(domain donationtypes.Domain) DonationTypes {
	return DonationTypes{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		Donations:  domain.Donations,
	}
}

func ToListDomain(data []DonationTypes) []donationtypes.Domain {
	result := []donationtypes.Domain{}
	for _, domain := range data {
		result = append(result, domain.ToDomain())
	}
	return result
}
