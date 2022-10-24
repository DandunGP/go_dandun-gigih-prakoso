package entities

import "gorm.io/gorm"

type User struct {
	*gorm.Model

	Email    string `json:"email" form:"Email"`
	Password string `json:"password" form:"Password"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
