package controllers

import (
	"APIMiddleware/praktikum/config"
	"APIMiddleware/praktikum/middleware"
	"APIMiddleware/praktikum/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "login failed",
			"error":   err.Error(),
		})
	}

	userResponse := models.UserResponse{int(user.ID), user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"user":    userResponse,
	})
}
