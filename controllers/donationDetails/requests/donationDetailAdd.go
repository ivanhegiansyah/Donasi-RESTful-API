package requests

import donationdetails "finalproject-BE/business/donationDetails"

type DonationDetailAdd struct {
	Description string `json:"description"`
}

func (donationDetail *DonationDetailAdd) ToDomainAdd() donationdetails.Domain {
	return donationdetails.Domain{
		Description: donationDetail.Description,
	}
}
