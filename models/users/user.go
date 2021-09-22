package users

import (
	"finalproject-BE/models/donations"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint `gorm:"primarykey"`
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email" gorm:"unique"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Dob       string `json:"dob" form:"dob"`
	CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
	Donations []donations.Donation
}
