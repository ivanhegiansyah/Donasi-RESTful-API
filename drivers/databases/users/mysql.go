package users

import (
	"context"
	"errors"
	"finalproject-BE/business/users"

	// "finalproject-BE/drivers/databases/donations"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	rep.Conn.Raw("SELECT password FROM users WHERE email = ?", domain.Email).Scan(&domain.Password)

	result := rep.Conn.First(&user, "email = ?", domain.Email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	user = FromDomain(domain)

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) GetAllUser(ctx context.Context) ([]users.Domain, error) {
	var user []Users

	result := rep.Conn.Find(&user)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}

	return ToListDomain(user), nil
}

func (rep *MysqlUserRepository) GetDetailUser(ctx context.Context, id int) (users.Domain, error) {
	var user Users
	// var donation []donations.Donations
	// var detail []users.DetailDonate
	// result := rep.Conn.First(&user, id)
	
	result := rep.Conn.Preload("Transaction").Preload("Donation").Find(&user, id)
	// result := rep.Conn.Preload("Donation", func(db *gorm.DB) *gorm.DB {
    //     return db.Select("Id", "UserId", "DonationName", "Status", "ShortDescription", "GoalAmount", "ExpiredDate")
    // }).Find(&user, id)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	userDomain := user.ToDomain()
	// temp := rep.Conn.Model(&donation).Select("donation.id, donation.donation_name, donation.status, donation.short_description, donation.goal_amount, donation.expired_date").Joins("JOIN users on donations.user_id = users.id").Where("users.id = ?", id).Scan(&detail)
	return userDomain, nil
}

func (rep *MysqlUserRepository) UpdateUser(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user Users

	user = FromDomain(domain)

	result := rep.Conn.Save(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil

}

func (rep *MysqlUserRepository) DeleteUser(ctx context.Context, id int) error {
	var user Users

	result := rep.Conn.Delete(&user, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}