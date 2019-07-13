package main

import (
	"github.com/jessicapaz/desafio-stone/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route
	e.GET("/", handlers.HealthCheck)
	e.POST("/users", handlers.CreateUser)
	e.POST("/login", handlers.Login)

	// Start server
	e.Logger.Fatal(e.Start(":8966"))
}
