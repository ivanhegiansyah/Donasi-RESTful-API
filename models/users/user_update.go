package users

type UserUpdate struct {
	Id       int    `json:"id" form:"id" gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Dob      string `json:"dob" form:"dob"`
}