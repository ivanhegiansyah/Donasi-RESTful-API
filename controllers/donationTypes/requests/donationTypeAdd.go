package requests

import donationtypes "finalproject-BE/business/donationTypes"

type DonationTypeAdd struct {
	Name string `json:"name"`
}

func (donationType *DonationTypeAdd) ToDomainAdd() donationtypes.Domain {
	return donationtypes.Domain{
		Name: donationType.Name,
	}
}
