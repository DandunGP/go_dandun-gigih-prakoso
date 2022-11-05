package usecase

import (
	"weekly3/entities"
	"weekly3/middleware"
	"weekly3/repository"
)

type LoginUsecase interface {
	Token(userData entities.User) (token string, err error)
}

type loginUsecase struct {
	loginRepository repository.LoginRepository
}

func NewLoginUsecase(loginRepo repository.LoginRepository) *loginUsecase {
	return &loginUsecase{loginRepository: loginRepo}
}

func (s *loginUsecase) Token(userData entities.User) (token string, err error) {

	data, err := s.loginRepository.GetUser(userData.Username)
	if err != nil {
		return token, err
	}

	if data.Password != userData.Password {
		return token, err
	}

	token, err = middleware.CreateToken(int(data.ID), data.Username)
	if err != nil {
		return "error", err
	}

	return token, nil
}
