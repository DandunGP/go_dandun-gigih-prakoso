package repository

import (
	"weekly3/entities"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAllUsers(data []entities.User) ([]entities.User, error)
	CreateUser(data entities.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetAllUsers(data []entities.User) ([]entities.User, error) {
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *userRepository) CreateUser(data entities.User) error {
	return r.db.Create(&data).Error
}
