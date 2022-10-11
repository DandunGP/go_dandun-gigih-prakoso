package database

import (
	"APIMiddleware/praktikum/config"
	"APIMiddleware/praktikum/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
