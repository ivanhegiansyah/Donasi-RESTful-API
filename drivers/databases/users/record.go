package users

import (
	"finalproject-BE/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id        int 			`gorm:"primaryKey"`
	Name      string
	Email     string 		` gorm:"unique"`
	Password  string
	Phone     string
	Dob       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Phone: user.Phone,
		Dob: user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id: domain.Id,
		Name: domain.Name,
		Email: domain.Email,
		Password: domain.Password,
		Phone: domain.Phone,
		Dob: domain.Dob,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}