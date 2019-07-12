package main

import (
    "github.com/jessicapaz/desafio-stone/handlers"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", handlers.HealthCheck)

    e.Logger.Fatal(e.Start(":8966"))
}
