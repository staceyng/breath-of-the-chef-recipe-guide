package main

import (
	"breath-of-the-chef-recipe-guide/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.HelloWorld)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
