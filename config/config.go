package config

import (
	"finalproject-BE/models/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:ivan112233@tcp(127.0.0.1:3306)/donasi?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal terhubung ke server database")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&users.User{})
}
