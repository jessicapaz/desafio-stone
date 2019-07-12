package handlers

import (
    "net/http"
    "github.com/jessicapaz/desafio-stone/renderings"
    "github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
    resp := renderings.HealthCheckResponse{
        Message: "OK",
    }
    return c.JSON(http.StatusOK, resp)
}
