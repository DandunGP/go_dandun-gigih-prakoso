package usecase

import (
	"weekly3/entities"
	"weekly3/repository"
)

type UserUsecase interface {
	Find(userData []entities.User) ([]entities.User, error)
	Create(userData entities.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) Create(userData entities.User) error {

	if err := s.userRepository.CreateUser(userData); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) Find(userData []entities.User) ([]entities.User, error) {

	if data, err := s.userRepository.GetAllUsers(userData); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
