package main

import (
	"testing/RESTfulAPI/config"
	m "testing/RESTfulAPI/middleware"
	"testing/RESTfulAPI/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
