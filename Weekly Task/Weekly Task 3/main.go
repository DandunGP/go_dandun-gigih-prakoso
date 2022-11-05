package main

import (
	"weekly3/config"
	m "weekly3/middleware"
	"weekly3/routes"

	"github.com/labstack/echo"
)

func main() {
	db := config.InitDB()
	app := echo.New()
	routes.New(app, db)
	m.LogMiddleware(app)
	app.Logger.Fatal(app.Start(":8000"))
}
