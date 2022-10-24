package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user []User

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			user = append(user, users[i])
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "delete success",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	id, _ := strconv.Atoi(c.Param("id"))

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users[i].Name = name
			users[i].Email = email
			users[i].Password = password
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Update success",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	name := c.FormValue("Name")
	email := c.FormValue("Email")
	password := c.FormValue("Password")

	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	user.Name = name
	user.Email = email
	user.Password = password
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
