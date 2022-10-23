package usecase

import (
	"belajar-go-echo/entities"
	"belajar-go-echo/middleware"
	"belajar-go-echo/repository"
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

	data, err := s.loginRepository.GetUser(userData.Email)
	if err != nil {
		return token, err
	}

	if data.Password != userData.Password {
		return token, err
	}

	token, err = middleware.CreateToken(int(data.ID), data.Email)
	if err != nil {
		return "error", err
	}

	return token, nil
}
