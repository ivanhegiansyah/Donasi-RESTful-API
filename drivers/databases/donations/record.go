package donations

import (
	"finalproject-BE/business/donations"
	donationdetails "finalproject-BE/drivers/databases/donationDetails"
	"finalproject-BE/drivers/databases/transactions"
	"time"

	"gorm.io/gorm"
)

type Donations struct {
	Id               int `gorm:"primaryKey"`
	UserId           int
	DonationTypeId   int
	DonationName     string
	Status           string
	ShortDescription string
	GoalAmount       int
	CurrentAmount    int
	ExpiredDate      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt                  `gorm:"index"`
	DonationDetails  donationdetails.DonationDetails `gorm:"foreignKey:DonationId;references:Id"`
	Transactions     []transactions.Transactions     `gorm:"many2many:donations_transactions;"`
}

func (donation *Donations) ToDomain() donations.Domain {
	return donations.Domain{
		Id:               donation.Id,
		UserId:           donation.UserId,
		DonationTypeId:   donation.DonationTypeId,
		DonationName:     donation.DonationName,
		Status:           donation.Status,
		ShortDescription: donation.ShortDescription,
		GoalAmount:       donation.GoalAmount,
		CurrentAmount:    donation.CurrentAmount,
		ExpiredDate:      donation.ExpiredDate,
		CreatedAt:        donation.CreatedAt,
		UpdatedAt:        donation.UpdatedAt,
	}
}

func FromDomain(domain donations.Domain) Donations {
	return Donations{
		Id:               domain.Id,
		UserId:           domain.UserId,
		DonationTypeId:   domain.DonationTypeId,
		DonationName:     domain.DonationName,
		Status:           "Active",
		ShortDescription: domain.ShortDescription,
		GoalAmount:       domain.GoalAmount,
		CurrentAmount:    0,
		ExpiredDate:      domain.ExpiredDate,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}

func ToListDomain(data []Donations) []donations.Domain {
	result := []donations.Domain{}
	for _, domain := range data {
		result = append(result, domain.ToDomain())
	}
	return result
}
