package main

import (
	"API/problem2/models"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {

	config := Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "crud_go",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func GetUsersController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	var users models.User

	if err := DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
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

	if err := DB.Save(&user).Error; err != nil {
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

	if err := DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Record not found!",
		})
	}

	DB.Delete(&users, c.Param("id"))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success detele user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	var users models.User

	if err := DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
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

	DB.Model(&users).Updates(input)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "update success",
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
