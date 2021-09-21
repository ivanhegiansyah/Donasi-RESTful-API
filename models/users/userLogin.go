package users

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
}