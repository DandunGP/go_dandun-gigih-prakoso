package controller

import (
	"belajar-go-echo/entities"
	"belajar-go-echo/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	var users []entities.User
	user, err := u.useCase.Find(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"data": user,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	var user entities.User
	c.Bind(&user)
	err := u.useCase.Create(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create new user",
		"data":    user,
	})
}
