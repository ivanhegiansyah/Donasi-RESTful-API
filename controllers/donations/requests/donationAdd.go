package requests

import (
	"finalproject-BE/business/donations"
)

type DonationAdd struct {
	UserId           int    `json:"userId"`
	DonationTypeId   int    `json:"donationTypeId"`
	DonationName     string `json:"donationName"`
	Status           string `json:"status"`
	ShortDescription string `json:"shortDescription"`
	GoalAmount       int    `json:"goalAmount"`
	ExpiredDate      string `json:"expiredDate"`
}

func (donation *DonationAdd) ToDomainAdd() donations.Domain {
	return donations.Domain{
		UserId:           donation.UserId,
		DonationTypeId:   donation.DonationTypeId,
		DonationName:     donation.DonationName,
		Status:           donation.Status,
		ShortDescription: donation.ShortDescription,
		GoalAmount:       donation.GoalAmount,
		ExpiredDate:      donation.ExpiredDate,
	}
}
