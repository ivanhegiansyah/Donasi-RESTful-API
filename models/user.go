package models

import (
	_ "time"

	_ "gorm.io/gorm"
)

type User struct {
	// gorm.Model
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Phone      string `json:"phone" form:"phone"`
	Dob        string `json:"dob" form:"dob"`
	// Created_at time.Time
	// Updated_at time.Time
}
