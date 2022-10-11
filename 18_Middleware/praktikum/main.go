package main

import (
	"APIMiddleware/praktikum/config"
	m "APIMiddleware/praktikum/middleware"
	"APIMiddleware/praktikum/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
