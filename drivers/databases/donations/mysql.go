package donations

import (
	"context"
	"errors"
	"finalproject-BE/business/donations"
	donationdetails "finalproject-BE/drivers/databases/donationDetails"
	"fmt"

	"gorm.io/gorm"
)

type MysqlDonationRepository struct {
	Conn *gorm.DB
}

func NewMysqlDonationRepository(conn *gorm.DB) donations.Repository {
	return &MysqlDonationRepository{
		Conn: conn,
	}
}

func (rep *MysqlDonationRepository) AddDonation(ctx context.Context, domain donations.Domain) (donations.Domain, error) {
	var donation Donations
	
	donation = FromDomain(domain)

	result := rep.Conn.Create(&donation)

	if result.Error != nil {
		return donations.Domain{}, result.Error
	}

	return donation.ToDomain(), nil
}

func (rep *MysqlDonationRepository) GetAllDonation(ctx context.Context) ([]donations.Domain, error) {
	var donation []Donations

	result := rep.Conn.Find(&donation)

	if result.Error != nil {
		return []donations.Domain{}, result.Error
	}

	return ToListDomain(donation), nil
}

func (rep *MysqlDonationRepository) GetDetailDonation(ctx context.Context, id int) (donations.Domain, error) {
	var donation Donations
	//sebelum find
	result := rep.Conn.Preload("DonationDetails").First(&donation, id)
	
	

	if result.Error != nil {
		return donations.Domain{}, result.Error
	}

	return donation.ToDomain(), nil
}

func (rep *MysqlDonationRepository) UpdateDonation(ctx context.Context, domain donations.Domain) (donations.Domain, error) {
	var donation Donations

	donation = FromDomain(domain)

	result := rep.Conn.Save(&donation)

	if result.Error != nil {
		return donations.Domain{}, result.Error
	}

	return donation.ToDomain(), nil

}

func (rep *MysqlDonationRepository) DeleteDonation(ctx context.Context, id int) error {
	var donation Donations
	// var donationdetails donationdetails.DonationDetails
	// var idx, idx2 int

	// rep.Conn.Raw("SELECT id FROM donations").Scan(&idx)

	// rep.Conn.Raw("SELECT id FROM donation_details  WHERE donation_id = ?", idx).Scan(&idx2)
	// fmt.Println(idx2)
	// rep.Conn.Delete(&donationdetails, idx2)
	

	result := rep.Conn.Delete(&donation, id)
	
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (d *Donations) BeforeDelete(tx *gorm.DB) (err error) {
	donationDetails := donationdetails.DonationDetails{}
	var temp int

	tx.Raw("SELECT donation_id FROM donation_details").Scan(&temp)

	fmt.Println(temp)
	fmt.Println(d.Id)
	if d.Id == temp {
		tx.Delete(&donationDetails, d.Id)
	}

	return
}