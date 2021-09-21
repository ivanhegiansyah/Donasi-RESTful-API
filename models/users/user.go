package users

import (
	"finalproject-BE/models/donations"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" gorm:"unique"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Dob       string `json:"dob" form:"dob"`
	Donations []donations.Donation
}
