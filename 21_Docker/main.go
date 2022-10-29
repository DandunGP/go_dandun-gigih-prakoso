package main

import (
	"testdocker/config"
	m "testdocker/middleware"
	"testdocker/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
