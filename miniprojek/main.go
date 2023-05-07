package main

import (
	"miniprojek/config"
	"miniprojek/middlewares"
	"miniprojek/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middlewares.Logmiddleware(e)

	routes.New(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
