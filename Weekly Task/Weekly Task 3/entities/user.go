package entities

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"Name"`
	Username string `json:"username" form:"Username"`
	Password string `json:"password" form:"Password"`
}

type UserResponse struct {
	ID       int    `json:"id" form:"name"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Token    string `json:"token" form:"token"`
}
