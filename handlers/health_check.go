package handlers

import (
	"github.com/jessicapaz/desafio-stone/renderings"
	"github.com/labstack/echo"
	"net/http"
)

// HealthCheck returns a successful message if the api is working
func HealthCheck(c echo.Context) error {
	resp := renderings.HealthCheckResponse{
		Message: "Everything is ok",
	}
	return c.JSON(http.StatusOK, resp)
}
