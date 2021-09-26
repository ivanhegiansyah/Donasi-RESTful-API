package responses

import (
	donationdetails "finalproject-BE/business/donationDetails"
	"time"
)

type DonationDetailResponse struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain donationdetails.Domain) DonationDetailResponse {
	return DonationDetailResponse{
		ID:          domain.Id,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
