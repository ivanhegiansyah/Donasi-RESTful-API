package responses

import (
	"finalproject-BE/business/donations"
	donationdetails "finalproject-BE/drivers/databases/donationDetails"
)

type DonationResponse struct {
	Id               int    `json:"id"`
	UserId           int    `json:"userId"`
	DonationTypeId   int    `json:"donationTypeId"`
	DonationName     string `json:"donationName"`
	Status           string `json:"status"`
	ShortDescription string `json:"shortDescription"`
	GoalAmount       int    `json:"goalAmount"`
	ExpiredDate      string `json:"expiredDate"`
}

type DonationDetailResponse struct {
	DonationResponses DonationResponse
	CurrentAmount     int                             `json:"currentAmount"`
	DonationDetail    donationdetails.DonationDetails `json:"donationDetail"`
}

func FromDomainDetail(domain donations.Domain) DonationDetailResponse {
	return DonationDetailResponse{
		DonationResponses: DonationResponse{
			Id:               domain.Id,
			UserId:           domain.UserId,
			DonationTypeId:   domain.DonationTypeId,
			DonationName:     domain.DonationName,
			Status:           domain.Status,
			ShortDescription: domain.ShortDescription,
			ExpiredDate:      domain.ExpiredDate,
			GoalAmount:       domain.GoalAmount,
		},
		CurrentAmount:  domain.CurrentAmount,
		DonationDetail: domain.DonationDetail,
	}
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
		ExpiredDate:      domain.ExpiredDate,
	}
}

func FromDonationListDomain(domain []donations.Domain) []DonationResponse {
	var response []DonationResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}
