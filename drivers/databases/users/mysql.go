package users

import (
	"context"
	"finalproject-BE/business/users"

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

	result := rep.Conn.First(&user, id)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

