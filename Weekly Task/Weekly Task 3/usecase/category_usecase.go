package usecase

import (
	"weekly3/entities"
	"weekly3/repository"
)

type CategoryUsecase interface {
	Find(categoryData []entities.Category) ([]entities.Category, error)
	Create(categoryData entities.Category) error
}

type categoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepository: categoryRepo}
}

func (s *categoryUsecase) Create(categoryData entities.Category) error {

	if err := s.categoryRepository.CreateCategory(categoryData); err != nil {
		return err
	}

	return nil
}

func (s *categoryUsecase) Find(categoryData []entities.Category) ([]entities.Category, error) {

	if data, err := s.categoryRepository.GetAllCategorys(categoryData); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
