package main

import (
	"github.com/jessicapaz/desafio-stone/config"
	"github.com/jessicapaz/desafio-stone/handlers"
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/services"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := config.GetDB()
	userH := handlers.NewUserHandler(models.NewUserModel(db))
	loginH := handlers.NewLoginHandler(services.NewLoginService())
	// Route
	e.GET("/", handlers.HealthCheck)
	e.POST("/users", userH.CreateUser)
	e.POST("/login", loginH.Login)

	// Start server
	e.Logger.Fatal(e.Start(":8966"))
}
