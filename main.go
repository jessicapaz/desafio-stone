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
	handler := handlers.NewHandler(
		models.NewUserModel(db),
		services.NewLoginService(),
		models.NewInvoiceModel(db),
	)

	// Route
	e.GET("/", handlers.HealthCheck)
	e.POST("/users", handler.CreateUser)
	e.POST("/login", handler.Login)

	// Start server
	e.Logger.Fatal(e.Start(":8966"))
}
