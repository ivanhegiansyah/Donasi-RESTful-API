package responses

import (
	donationtypes "finalproject-BE/business/donationTypes"
	"finalproject-BE/drivers/databases/donations"
	"time"
)

type DonationTypeResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DonationTypeDetailResponse struct {
	DonationTypeResponses DonationTypeResponse
	Donations             []donations.Donations `json:"donations"`
}

func FromDomainDetail(domain donationtypes.Domain) DonationTypeDetailResponse {
	return DonationTypeDetailResponse{
		DonationTypeResponses: DonationTypeResponse{
			Id:   domain.Id,
			Name: domain.Name,
		},
		Donations: domain.Donations,
	}
}

func FromDomain(domain donationtypes.Domain) DonationTypeResponse {
	return DonationTypeResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromListDomain(domain []donationtypes.Domain) []DonationTypeResponse {
	var response []DonationTypeResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}
