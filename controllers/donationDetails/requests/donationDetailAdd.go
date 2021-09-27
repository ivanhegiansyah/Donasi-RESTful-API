package requests

import donationdetails "finalproject-BE/business/donationDetails"

type DonationDetailAdd struct {
	DonationId int `json:"donationId"`
	Description string `json:"description"`
}

func (donationDetail *DonationDetailAdd) ToDomainAdd() donationdetails.Domain {
	return donationdetails.Domain{
		DonationId: donationDetail.DonationId,
		Description: donationDetail.Description,
	}
}
