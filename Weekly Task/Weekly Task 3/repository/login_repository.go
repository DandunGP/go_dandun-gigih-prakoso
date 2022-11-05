package repository

import (
	"weekly3/entities"

	"github.com/jinzhu/gorm"
)

type LoginRepository interface {
	GetUser(username string) (entities.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *loginRepository {
	return &loginRepository{db}
}

func (r *loginRepository) GetUser(username string) (entities.User, error) {
	var users entities.User
	if err := r.db.Where("username = ?", username).First(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
