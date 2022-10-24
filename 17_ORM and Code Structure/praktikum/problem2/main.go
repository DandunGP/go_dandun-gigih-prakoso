package main

import (
	"API/problem2/config"
	"API/problem2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
