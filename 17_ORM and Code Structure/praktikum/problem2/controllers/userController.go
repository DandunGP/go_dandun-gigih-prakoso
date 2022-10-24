package controllers

import (
	"API/problem2/config"
	"API/problem2/lib/database"
	"API/problem2/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success get all users",
		"users":  users,
	})
}

func GetUserController(c echo.Context) error {
	var users models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	var user models.User
	c.Bind(&user)

	user.Name = name
	user.Email = email
	user.Password = password

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	var users models.User

	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	config.DB.Delete(&users, id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success detele user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var users models.User

	if err := config.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	var input models.User

	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	input.Name = name
	input.Email = email
	input.Password = password

	config.DB.Model(&users).Updates(input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}
