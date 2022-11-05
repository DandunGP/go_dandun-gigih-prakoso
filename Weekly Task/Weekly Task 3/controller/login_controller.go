package controller

import (
	"net/http"
	"weekly3/entities"
	"weekly3/usecase"

	"github.com/labstack/echo"
)

type LoginController interface{}

type loginController struct {
	useCase usecase.LoginUsecase
}

func NewLoginController(loginUsecase usecase.LoginUsecase) *loginController {
	return &loginController{
		loginUsecase,
	}
}

func (l *loginController) GetUser(c echo.Context) error {
	user := entities.User{}
	c.Bind(&user)

	token, err := l.useCase.Token(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"user":    token,
	})
}
