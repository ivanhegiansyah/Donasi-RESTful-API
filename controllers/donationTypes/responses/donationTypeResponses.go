package responses

import (
	donationtypes "finalproject-BE/business/donationTypes"
	"time"
)

type DonationTypeResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain donationtypes.Domain) DonationTypeResponse {
	return DonationTypeResponse{
		Id:        domain.Id,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
