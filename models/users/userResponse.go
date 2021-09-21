package users

type UserResponse struct {
	Id       uint    `json:"id" form:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Token string `json:"token" form:"token"`
}