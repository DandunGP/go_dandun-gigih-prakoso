package repository

import (
	"belajar-go-echo/entities"

	"gorm.io/gorm"
)

type LoginRepository interface {
	GetUser(email string) (entities.User, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *loginRepository {
	return &loginRepository{db}
}

func (r *loginRepository) GetUser(email string) (entities.User, error) {
	var users entities.User
	// query := fmt.Sprintf("SELECT * FROM users WHERE email = '%s'", email)
	// err = r.db.Select().Where()
	if err := r.db.Where("email = ?", email).First(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
