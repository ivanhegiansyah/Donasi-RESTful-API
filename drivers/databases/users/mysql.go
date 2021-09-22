package users

import (
	"context"
	"finalproject-BE/business/users"
	"finalproject-BE/controllers/users/requests"

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

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Register(ctx context.Context, userRegister requests.UserRegister) (users.Domain, error) {
	var user Users
	
	user.Name = userRegister.Name
	user.Email = userRegister.Email
	user.Password = userRegister.Password
	user.Phone = userRegister.Phone
	user.Dob = userRegister.Dob

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
