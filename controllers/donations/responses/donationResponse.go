package responses

import (
	"finalproject-BE/business/donations"
	"time"
)

type DonationResponse struct {
	Id               int       `json:"id"`
	UserId           int       `json:"userId"`
	DonationTypeId   int       `json:"donationTypeId"`
	DonationName     string    `json:"donationName"`
	Status           string    `json:"status"`
	ShortDescription string    `json:"shortDescription"`
	GoalAmount       int       `json:"goalAmount"`
	CurrentAmount    int       `json:"currentAmount"`
	ExpiredDate      string    `json:"expiredDate"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func FromDomain(domain donations.Domain) DonationResponse {
	return DonationResponse{
		Id:               domain.Id,
		UserId:           domain.UserId,
		DonationTypeId:   domain.DonationTypeId,
		DonationName:     domain.DonationName,
		Status:           domain.Status,
		ShortDescription: domain.ShortDescription,
		GoalAmount:       domain.GoalAmount,
		CurrentAmount:    domain.CurrentAmount,
		ExpiredDate:      domain.ExpiredDate,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
	}
}
