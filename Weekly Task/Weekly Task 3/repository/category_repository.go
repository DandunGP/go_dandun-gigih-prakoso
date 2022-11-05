package repository

import (
	"weekly3/entities"

	"github.com/jinzhu/gorm"
)

type CategoryRepository interface {
	GetAllCategorys(data []entities.Category) ([]entities.Category, error)
	CreateCategory(data entities.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategorys(data []entities.Category) ([]entities.Category, error) {
	if err := r.db.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *categoryRepository) CreateCategory(data entities.Category) error {
	return r.db.Create(&data).Error
}
