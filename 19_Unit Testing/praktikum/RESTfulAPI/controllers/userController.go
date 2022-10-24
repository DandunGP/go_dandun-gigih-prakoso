package controllers

import (
	"net/http"
	"strconv"
	"testing/RESTfulAPI/config"
	"testing/RESTfulAPI/lib/database"
	"testing/RESTfulAPI/models"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all users",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", id).First(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"users":   users,
	})
}

func CreateUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	user.Name = name
	user.Email = email
	user.Password = password

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	var input models.User

	input.Name = name
	input.Email = email
	input.Password = password

	if err := config.DB.Model(&users).Where("id = ?", id).Updates(input).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func DeleteUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Record not found!")
	}

	config.DB.Delete(&users, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}
