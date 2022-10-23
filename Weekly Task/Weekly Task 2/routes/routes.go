package routes

import (
	"weekly2/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/items", controllers.GetItemsController)
	e.GET("/items/:id", controllers.GetItemController)
	e.POST("/items", controllers.CreateItemController)
	e.PUT("/items/:id", controllers.UpdateItemController)
	e.DELETE("/items/:id", controllers.DeleteItemController)
	e.GET("/items/category/:category_id", controllers.GetItemByCategory)
	e.GET("/items", controllers.GetItemByName)
	return e
}

/*
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
*/
