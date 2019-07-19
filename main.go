package main

import (
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jessicapaz/desafio-stone/config"
	"github.com/jessicapaz/desafio-stone/handlers"
	"github.com/jessicapaz/desafio-stone/models"
	"github.com/jessicapaz/desafio-stone/services"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
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

	r := e.Group("/invoices")
	r.Use(middleware.JWT([]byte(os.Getenv("TOKEN_PASSWORD"))))
	r.POST("", handler.CreateInvoice)
	r.GET("", handler.ListInvoice)
	r.DELETE("/:id", handler.DeactivateInvoice)
	r.GET("/:id", handler.RetrieveInvoice)

	// Start server
	e.Logger.Fatal(e.Start(":8966"))
}
