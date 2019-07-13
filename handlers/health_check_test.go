package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	e := echo.New()
	e.GET("/", HealthCheck)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	e.ServeHTTP(rec, req)
	resp := rec.Result()
	if resp.StatusCode != http.StatusOK {
		t.Error("unexpected status code: ", resp.Status)
	}
}
