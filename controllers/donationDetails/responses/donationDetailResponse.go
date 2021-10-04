package responses

import (
	donationdetails "finalproject-BE/business/donationDetails"
	"time"
)

type DonationDetailResponse struct {
	ID          int       `json:"id"`
	DonationId  int       `json:"donationId"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain donationdetails.Domain) DonationDetailResponse {
	return DonationDetailResponse{
		ID:          domain.Id,
		DonationId:  domain.DonationId,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
