package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/constant"
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	loginRepository := repository.NewLoginRepository(db)

	userService := usecase.NewUserUsecase(userRepository)
	loginService := usecase.NewLoginUsecase(loginRepository)

	userController := controller.NewUserController(userService)
	loginController := controller.NewLoginController(loginService)

	e.POST("/login", loginController.GetUser)

	eJwtAuth := e.Group("auth")
	eJwtAuth.Use(mid.JWT([]byte(constant.SECRET_KEY)))
	eJwtAuth.GET("/users", userController.GetAllUsers)
	eJwtAuth.POST("/users", userController.CreateUser)
}

func main() {
	db := config.InitDB()

	app := echo.New()

	NewRoute(app, db)

	app.Logger.Fatal(app.Start(":8000"))
}
