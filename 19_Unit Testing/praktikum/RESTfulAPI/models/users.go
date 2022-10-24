package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID    int    `json:"id" form:"name"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
